package abi

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"
)

var endpoint = "http://127.0.0.1:8502"
var wsEndpoint = "ws://127.0.0.1:8547"
var validatorSetAddr = common.HexToAddress("0x0000000000000000000000000000000000001000")
var crosschainAddr = common.HexToAddress("0x0000000000000000000000000000000000001004")
var systemRewardAddr = common.HexToAddress("0x0000000000000000000000000000000000001002")
var account, _ = fromHexKey("9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3")
var receiveAccount = common.HexToAddress("0x27d92f736324e6d9f85d37a27a23aaabe7162168")
var validatorSetABI, _ = abi.JSON(strings.NewReader(ValidatorABI))

func TestSimulateUpdateValidatorContract(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)

	instance, err := NewValidator(validatorSetAddr, client)
	assert.NoError(t, err)
	nonce, err := client.PendingNonceAt(context.Background(), account.addr)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)
	auth := bind.NewKeyedTransactor(account.key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	sequence, err := instance.Sequence(nil)
	assert.NoError(t, err)

	wsclient, err := ethclient.Dial(wsEndpoint)
	assert.NoError(t, err)
	wsInstance, err := NewValidator(validatorSetAddr, wsclient)
	assert.NoError(t, err)
	sink := make(chan *ValidatorValidatorSetUpdated)
	subs, err := wsInstance.WatchValidatorSetUpdated(nil, sink)
	assert.NoError(t, err)
	defer subs.Unsubscribe()

	bz, err := instance.InitValidatorSetBytes(nil)
	assert.NoError(t, err)

	data, err := validatorSetABI.Pack("update", bz[1+68:],
		common.Hex2Bytes("12"), uint64(1), sequence+1)
	assert.NoError(t, err)

	msg := ethereum.CallMsg{
		auth.From,
		&validatorSetAddr,
		auth.GasLimit,
		auth.GasPrice,
		auth.Value,
		data,
	}

	res, err := client.CallContract(context.Background(), msg, nil)
	assert.NoError(t, err)
	fmt.Println(string(res))
}

// do update
func TestUpdateValidatorContract(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)

	instance, err := NewValidator(validatorSetAddr, client)
	assert.NoError(t, err)
	nonce, err := client.PendingNonceAt(context.Background(), account.addr)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)
	auth := bind.NewKeyedTransactor(account.key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	sequence, err := instance.Sequence(nil)
	assert.NoError(t, err)

	wsclient, err := ethclient.Dial(wsEndpoint)
	assert.NoError(t, err)
	wsInstance, err := NewValidator(validatorSetAddr, wsclient)
	assert.NoError(t, err)
	sink := make(chan *ValidatorValidatorSetUpdated)
	subs, err := wsInstance.WatchValidatorSetUpdated(nil, sink)
	assert.NoError(t, err)
	defer subs.Unsubscribe()
	time.Sleep(1 * time.Second)

	bz, err := instance.InitValidatorSetBytes(nil)
	assert.NoError(t, err)
	tx, e := instance.Update(auth, append([]byte{0x00},bz[1:]...), common.Hex2Bytes("12"), 0, sequence+1)

	assert.NoError(t, e)
	time.Sleep(4 * time.Second)
	fmt.Println("wait event")
	//<-sink
	r, err := client.TransactionReceipt(context.Background(), tx.Hash())
	assert.NoError(t, err)
	fmt.Printf("status %d\n",r.Status)
	fmt.Printf("gas used %d\n", r.GasUsed)
	fmt.Printf("gas price %v\n", tx.GasPrice().String())
	fmt.Printf("success: %v\n", r.Status == 1)

	vs,_:=instance.GetValidators(nil);
	fmt.Println(len(vs))
}

func TestGetBalanceOfValidators(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)
	instance, err := NewValidator(validatorSetAddr, client)
	assert.NoError(t, err)
	vs, err := instance.GetValidators(nil)
	assert.NoError(t, err)
	for idx, v := range vs {
		income, err := instance.GetIncoming(nil, v)
		assert.NoError(t, err)
		validator, err := instance.CurrentValidatorSet(nil, big.NewInt(int64(idx)))
		assert.NoError(t, err)
		b, err := client.BalanceAt(context.Background(), validator.FeeAddress, nil)
		assert.NoError(t, err)
		fmt.Printf("%s has %v in contract while %v in account\n", v.String(), income.String(), b.String())
	}
	v, err := client.BalanceAt(context.Background(), crosschainAddr, nil)
	assert.NoError(t, err)
	fmt.Printf("cross chain have %v\n", v.String())

	v, err = client.BalanceAt(context.Background(), systemRewardAddr, nil)
	assert.NoError(t, err)
	fmt.Printf("system reward have %v\n", v.String())

	v, err = client.BalanceAt(context.Background(), account.addr, nil)
	assert.NoError(t, err)
	fmt.Printf("myaccount have %v\n", v.String())

	//0x7CFd964E035e1435D78C58EF87117dc3f5039958 has 4250589843750000000 in contract while 0 in account
	//0xA16e8d967D3959F13bF4c5E24f15C00B9D0887D9 has 4068750000000000000 in contract while 39376 in account
	//0x7698c2e4582D6174ED31171A7fe4A5641d4A98d2 has 4200000000000000000 in contract while 39376 in account
	// total 12519339843750000000
	//cross chain have 37913589843700000000
	//system reward have 1119245312550273418
	//myaccount have 906447824999999647830

	//gas used 107925
	//gas price 500000000000
	// total  53962500000000000

	//0x7CFd964E035e1435D78C58EF87117dc3f5039958 has 0 in contract while 0 in account
	//0xA16e8d967D3959F13bF4c5E24f15C00B9D0887D9 has 53962500000000000 in contract while 39376 in account
	//0x7698c2e4582D6174ED31171A7fe4A5641d4A98d2 has 0 in contract while 39376 in account
	//cross chain have 50432929687400000000   diff 12519339843700000000
	//											   12519339843750000000
	// 											    50000000 goes to system reward
	//system reward have 1019245312600273418    diff   99999999950000000
	//myaccount have 906493862499999647830   diff 46037500000000000  used 53962500000000000  46037500000000000
}

func TestTmp(t *testing.T) {

	fmt.Println(time.Unix(567818489952, 0))
}

func TestTransfer(t *testing.T) {
	for i := 0; i < 20; i++ {
		client, err := ethclient.Dial(endpoint)
		assert.NoError(t, err)
		balanceBefore, err := client.BalanceAt(context.Background(), receiveAccount, nil)
		assert.NoError(t, err)
		tx, err := sendEther(client, account, receiveAccount, big.NewInt(params.Ether), false)
		assert.NoError(t, err)
		time.Sleep(3 * time.Second)
		r, err := client.TransactionReceipt(context.Background(), tx)
		fmt.Println(err)
		assert.NoError(t, err)
		assert.Equal(t, r.Status, uint64(1))
		balanceAfter, err := client.BalanceAt(context.Background(), receiveAccount, nil)
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(0).Sub(balanceAfter, balanceBefore).Cmp(big.NewInt(params.Ether)), 0)
	}
}

func TestTransferWithHugeFee(t *testing.T) {
	for i := 0; i < 6; i++ {
		client, err := ethclient.Dial(endpoint)
		assert.NoError(t, err)
		balanceBefore, err := client.BalanceAt(context.Background(), receiveAccount, nil)
		assert.NoError(t, err)
		tx, err := sendEther(client, account, receiveAccount, big.NewInt(params.Ether), true)
		assert.NoError(t, err)
		time.Sleep(2 * time.Second)
		r, err := client.TransactionReceipt(context.Background(), tx)
		assert.NoError(t, err)
		assert.Equal(t, r.Status, uint64(1))
		balanceAfter, err := client.BalanceAt(context.Background(), receiveAccount, nil)
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(0).Sub(balanceAfter, balanceBefore).Cmp(big.NewInt(params.Ether)), 0)
	}
}

func TestBlockNum(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		panic(err)
	}

	b, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}
}

func TestBlocRecepient(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		panic(err)
	}

	b, err := client.BlockByNumber(context.Background(), big.NewInt(193))

	accounts := make(map[common.Address]bool, 0)
	assert.NoError(t, err)
	for i := 0; i < b.Transactions().Len(); i++ {
		r, err := client.TransactionReceipt(context.Background(), b.Transactions()[0].Hash())
		assert.NoError(t, err)
		fmt.Printf("status %d , gas used %d , gas price %d\n", r.Status, r.GasUsed, b.Transactions()[i].GasPrice())
		accounts[*b.Transactions()[i].To()] = true
		from, err := types.Sender(types.NewEIP155Signer(big.NewInt(714)), b.Transactions()[i])
		assert.NoError(t, err)
		accounts[from] = true
	}

	// get account
	beofreTotal := big.NewInt(0)
	for a := range accounts {
		b, err := client.BalanceAt(context.Background(), a, big.NewInt(192))
		beofreTotal = beofreTotal.Add(beofreTotal, b)
		assert.NoError(t, err)
	}

	afterTotal := big.NewInt(0)
	for a := range accounts {
		b, err := client.BalanceAt(context.Background(), a, big.NewInt(193))
		afterTotal = afterTotal.Add(afterTotal, b)
		assert.NoError(t, err)
	}

	fmt.Printf("before %s\n", beofreTotal.String())
	fmt.Printf("after %s\n", afterTotal.String())
	fmt.Printf("sub: %s\n", big.NewInt(0).Sub(beofreTotal, afterTotal))

}

func TestGetETH(t *testing.T) {
	fmt.Println(new(big.Float).Quo(new(big.Float).SetInt(math.MustParseBig256("99999999950000000")), new(big.Float).SetInt(big.NewInt(params.Ether))))
}

func TestTmp1(t *testing.T) {
	_, ok := context.Background().Value("as").([]byte)
	fmt.Println(ok)
}

func sendEther(client *ethclient.Client, fromEO ExtAcc, toAddr common.Address, value *big.Int, hugegasPrice bool) (common.Hash, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromEO.addr)
	if err != nil {
		return common.Hash{}, err
	}
	gasLimit := uint64(3e4)
	var gasPrice *big.Int
	if hugegasPrice {
		gasPrice = big.NewInt(params.Ether / 10000)
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
	}
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), fromEO.key)
	if err != nil {
		return common.Hash{}, err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, err
	}
	txhash := signedTx.Hash()
	log.Printf("tx hash, sendEther: %s\n", txhash.Hex())
	return txhash, nil
}

type ExtAcc struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

func fromHexKey(hexkey string) (ExtAcc, error) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return ExtAcc{}, err
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("publicKey is not of type *ecdsa.PublicKey")
		return ExtAcc{}, err
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return ExtAcc{key, addr}, nil
}

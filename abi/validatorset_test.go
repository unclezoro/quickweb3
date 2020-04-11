package abi

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
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
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"strings"
	"testing"
	"time"
)

//var endpoint = "http://127.0.0.1:8502"

var endpoint ="http://dex-qa-s1-bsc-dev-validator-alb-501442930.ap-northeast-1.elb.amazonaws.com:8545"


//var endpoint = "http://127.0.0.1:8545"

//var wsEndpoint = "ws://127.0.0.1:8547"
var wsEndpoint = "ws://dex-qa-s1-bsc-dev-validator-alb-501442930.ap-northeast-1.elb.amazonaws.com:8546"

var validatorSetAddr = common.HexToAddress("0x0000000000000000000000000000000000001000")
var crosschainAddr = common.HexToAddress("0x0000000000000000000000000000000000001004")
var systemRewardAddr = common.HexToAddress("0x0000000000000000000000000000000000001002")
var account, _ = fromHexKey("9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3")
var receiveAccount = common.HexToAddress("0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3")
var SystemAddress = common.HexToAddress("0xffffFFFfFFffffffffffffffFfFFFfffFFFfFFfE")
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

	bz, err := instance.INITVALIDATORSETBYTES(nil)
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

func TestA(t *testing.T) {
	bz, err := hex.DecodeString("0000010002080000000000000000")
	fmt.Println(len(bz))
	fmt.Println(err)
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

	bz, err := instance.INITVALIDATORSETBYTES(nil)
	assert.NoError(t, err)
	tx, e := instance.Update(auth, append([]byte{0x00}, bz[1:]...), common.Hex2Bytes("12"), 0, sequence+1)

	assert.NoError(t, e)
	time.Sleep(4 * time.Second)
	fmt.Println("wait event")
	//<-sink
	r, err := client.TransactionReceipt(context.Background(), tx.Hash())
	assert.NoError(t, err)
	fmt.Printf("status %d\n", r.Status)
	fmt.Printf("gas used %d\n", r.GasUsed)
	fmt.Printf("gas price %v\n", tx.GasPrice().String())
	fmt.Printf("success: %v\n", r.Status == 1)

	vs, _ := instance.GetValidators(nil)
	fmt.Println(len(vs))
}

func TestGetIncoming(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)
	instance, err := NewValidator(validatorSetAddr, client)
	assert.NoError(t, err)

	for h := int64(6016); h <= 6017; h++ {
		fmt.Printf("\nheiht %d \n", h)
		b, _ := client.BlockByNumber(context.Background(), big.NewInt(h))
		fmt.Printf("tx length at height %v %v \n", h, len(b.Transactions()))
		varls,_:=instance.GetValidators(&bind.CallOpts{BlockNumber: big.NewInt(h)})
		fmt.Printf("validators:")

		for _,v:=range varls{
			fmt.Printf("%s  ",v.String())
		}
		fmt.Printf("\n")

		for i := int64(0); i < 3; i++ {
			res, _ := instance.CurrentValidatorSet(&bind.CallOpts{BlockNumber: big.NewInt(h)}, big.NewInt(i))
			fmt.Printf("validator %s income %s\n", res.ConsensusAddress.String(), res.Incoming.String())
		}
		balance, _ := client.BalanceAt(context.Background(), common.HexToAddress("0x9fB29AAc15b9A4B7F17c3385939b007540f4d791"), big.NewInt(h))
		fmt.Printf("balance 0x9fB29AAc15b9A4B7F17c3385939b007540f4d791 %v\n", balance.String())
		v, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h))
		fmt.Printf("balance system reward %v\n", v.String())

		r, _ := client.BalanceAt(context.Background(), SystemAddress, big.NewInt(h))
		fmt.Printf("balance system account %v\n", r.String())

		k, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h))
		fmt.Printf("balance validatorset account %v\n", k.String())
	}

	b, _ := client.BlockByNumber(context.Background(), big.NewInt(6017))
	fmt.Println(b.Hash().String())

	r, _ := client.TransactionReceipt(context.Background(), b.Transactions()[0].Hash())
	fmt.Println(r.GasUsed)

	b1, _ := client.BlockByNumber(context.Background(), big.NewInt(6018))
	fmt.Println(b.Transactions()[0].To().String())

	r1, _ := client.TransactionReceipt(context.Background(), b1.Transactions()[0].Hash())
	fmt.Println(r1.GasUsed)
	end := uint64(6020)
	ite, err := instance.FilterValidatorDeposit(&bind.FilterOpts{Start: 6016, End: &end}, nil, nil)
	assert.NoError(t, err)
	for ite.Next() {
		fmt.Println("do have")
	}
}

func TestGetEvent(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	instance, err := NewValidator(validatorSetAddr, client)

	end:=uint64(258)
	ite,err:=instance.FilterValidatorDeposit(&bind.FilterOpts{Start:258,End:&end},nil,nil)
	assert.NoError(t,err)
	if ite.Next(){
		e:=ite.Event
		fmt.Printf("validator is %s\n",e.Validator.String())
		fmt.Printf("Amount is %s\n",e.Amount.String())
	}
}

func TestWatchEvent(t *testing.T) {
	wsclient, err := ethclient.Dial(wsEndpoint)
	assert.NoError(t, err)
	wsInstance, err := NewValidator(validatorSetAddr, wsclient)
	sink := make(chan *ValidatorValidatorDeposit)
	start := uint64(258)
	subs, err := wsInstance.WatchValidatorDeposit(&bind.WatchOpts{Start: &start}, sink,nil, nil)
	assert.NoError(t, err)
	defer subs.Unsubscribe()
	for s := range sink {
		fmt.Println(s.Raw.Data)
	}
}

func TestGetBlock(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)

	b, _ := client.BlockByNumber(context.Background(), big.NewInt(6017))
	signer,err:=recoverSigner(b.Header())
	assert.NoError(t,err)
	fmt.Println(signer.String())
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

	//client, err := ethclient.Dial(endpoint)
	//assert.NoError(t, err)

}
func TestLength(t *testing.T) {
	v := new(big.Int)
	fmt.Println(v.Cmp(common.Big0) == 0)

}

func TestTransfer(t *testing.T) {
	for i := 0; i < 2; i++ {
		client, err := ethclient.Dial(endpoint)
		assert.NoError(t, err)

		balanceBefore, err := client.BalanceAt(context.Background(), common.HexToAddress("0x9fB29AAc15b9A4B7F17c3385939b007540f4d791"), nil)
		fmt.Println(balanceBefore.String())
		assert.NoError(t, err)
		tx, err := sendEther(client, account, receiveAccount, big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(10000)), false)
		assert.NoError(t, err)
		time.Sleep(5 * time.Second)
		r, err := client.TransactionReceipt(context.Background(), tx)
		fmt.Println(r.BlockNumber)
		assert.NoError(t, err)
		assert.Equal(t, r.Status, uint64(1))
		balanceAfter, err := client.BalanceAt(context.Background(), receiveAccount, nil)
		fmt.Println(balanceAfter.String())
		assert.NoError(t, err)
		//assert.Equal(t, big.NewInt(0).Sub(balanceAfter, balanceBefore).Cmp(big.NewInt(params.Ether)), 0)
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

	b, err := client.BlockByNumber(context.Background(), big.NewInt(1))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b.Number().String())
		fmt.Println(b.Time())
		fmt.Println(b.Hash().String())
		fmt.Println(b.Transactions().Len())
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
	//log.Printf("tx hash, sendEther: %s\n", txhash.Hex())
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


type TransferInUnboundTokenEvent struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
}


func recoverSigner(header *types.Header) (common.Address, error) {
	signature := header.Extra[len(header.Extra)-65:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
	return signer, nil
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

func encodeSigHeader(w io.Writer, header *types.Header) {
	err := rlp.Encode(w, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-crypto.SignatureLength], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

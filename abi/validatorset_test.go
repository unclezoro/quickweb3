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
	geth "github.com/ethereum/go-ethereum/mobile"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"
)

var client *ethclient.Client
var endpoint = "http://127.0.0.1:8503"
var validatorSetAddr = common.HexToAddress("0x0000000000000000000000000000000000001000")
var account, _ = fromHexKey("9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3")
var receiveAccount = common.HexToAddress("0x27d92f736324e6d9f85d37a27a23aaabe7162168")
var validatorSetABI, _ = abi.JSON(strings.NewReader(ValidatorABI))

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

// simulate
func TestSimulateValidatorContract(t *testing.T) {
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)

	gasLimit := uint64(3e5)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)
	auth := bind.NewKeyedTransactor(account.key)
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	data, err := validatorSetABI.Pack("updateValidatorSet", common.Hex2Bytes("57d95bf4705b8636eef5060bbb51c470c6ad967557d95bf4705b8636eef5060bbb51c470c6ad967557d95bf4705b8636eef5060bbb51c470c6ad9675000000001000000069df37f10ecaf3c55c44bc2f2ca73610bb676d9469df37f10ecaf3c55c44bc2f2ca73610bb676d9469df37f10ecaf3c55c44bc2f2ca73610bb676d9400000000100000008ffd9b5276e4e9d18cb35e41b41fed69ed14fbcb8ffd9b5276e4e9d18cb35e41b41fed69ed14fbcb8ffd9b5276e4e9d18cb35e41b41fed69ed14fbcb0000000010000000"),
		common.Hex2Bytes("12"),  big.NewInt(1), big.NewInt(3))
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

	tx, e := instance.UpdateValidatorSet(auth, common.Hex2Bytes("57d95bf4705b8636eef5060bbb51c470c6ad967557d95bf4705b8636eef5060bbb51c470c6ad967557d95bf4705b8636eef5060bbb51c470c6ad9675000000001000000069df37f10ecaf3c55c44bc2f2ca73610bb676d9469df37f10ecaf3c55c44bc2f2ca73610bb676d9469df37f10ecaf3c55c44bc2f2ca73610bb676d9400000000100000008ffd9b5276e4e9d18cb35e41b41fed69ed14fbcb8ffd9b5276e4e9d18cb35e41b41fed69ed14fbcb8ffd9b5276e4e9d18cb35e41b41fed69ed14fbcb0000000010000000"),
		nil, big.NewInt(1), big.NewInt(3))

	assert.NoError(t, e)
	time.Sleep(2 * time.Second)
	r, err := client.TransactionReceipt(context.Background(), tx.Hash())
	assert.NoError(t, err)
	fmt.Printf("gas used %d\n", r.GasUsed)
	fmt.Printf("success: %v", r.Status == 1)
}


func TestTransfer(t *testing.T){
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)

	sendEther(account, receiveAccount, new(big.Int).Mul(big.NewInt(1), big.NewInt(params.GWei)))
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


func sendEther(fromEO ExtAcc, toAddr common.Address, value *big.Int) (common.Hash, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromEO.addr)
	if err != nil {
		return common.Hash{}, err
	}
	gasLimit := uint64(3e4)
	gasPrice, err := client.SuggestGasPrice(context.Background())
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
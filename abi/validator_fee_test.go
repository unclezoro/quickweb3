package abi

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestFee(t *testing.T) {
	go transferRoutin()
	client, _ := ethclient.Dial(endpoint)
	instance, _ := NewValidator(validatorSetAddr, client)
	headers := make(chan *types.Header, 100)
	wsclient, err := ethclient.Dial(wsEndpoint)
	assert.NoError(t, err)
	wsclient.SubscribeNewHead(context.Background(), headers)
	for header := range headers {
		fmt.Println("get")
		h := int64(header.Number.Uint64() - 7)
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(h)))
		assert.NoError(t, err)
		txs := block.Transactions()
		totalFee := big.NewInt(0)
		for _, tx := range txs {
			price := tx.GasPrice()
			recip, err := client.TransactionReceipt(context.Background(), tx.Hash())
			assert.NoError(t, err)
			totalFee = totalFee.Add(totalFee, price.Mul(price, big.NewInt(int64(recip.GasUsed))))
		}
		vA, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h))
		kA, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h))
		vb, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h-1))
		kb, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h-1))

		if big.NewInt(0).Add(vA, kA).Cmp(big.NewInt(0).Add(vb.Add(vb, kb), totalFee)) != 0 {
			panic(fmt.Sprintf("panic1 at hegiht %d", h))
		}
		totalIncome := big.NewInt(0)
		for i := int64(0); i < 3; i++ {
			res, err := instance.CurrentValidatorSet(&bind.CallOpts{BlockNumber: big.NewInt(h)}, big.NewInt(i))
			fmt.Println(err)
			totalIncome = totalIncome.Add(totalIncome, res.Incoming)
		}
		if totalIncome.Cmp(kA) != 0 {
			fmt.Println(totalIncome.String())
			fmt.Println(kA.String())
			panic(fmt.Sprintf("panic2 at hegiht %d", h))
		}
		fmt.Printf("check finish at height %d, txs %d, total income %s\n", h, len(txs), totalIncome.String())
		//fmt.Printf("check finish at height %d, txs %d, ",h,len(txs))

	}
}

func transferRoutin() {
	client, _ := ethclient.Dial(endpoint)
	for {
		sendEther(client, account, receiveAccount, big.NewInt(0).Mul(big.NewInt(params.GWei), big.NewInt(10000)), false)
		time.Sleep(1 * time.Second)
	}
}

func TestFee1(t *testing.T) {
	client, _ := ethclient.Dial(endpoint)
	instance, _ := NewValidator(validatorSetAddr, client)
	h := int64(43)
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(h)))
	assert.NoError(t, err)
	txs := block.Transactions()
	totalFee := big.NewInt(0)
	for _, tx := range txs {
		price := tx.GasPrice()
		recip, err := client.TransactionReceipt(context.Background(), tx.Hash())
		assert.NoError(t, err)
		totalFee = totalFee.Add(totalFee, price.Mul(price, big.NewInt(int64(recip.GasUsed))))
	}
	vA, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h))
	kA, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h))
	vb, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h-1))
	kb, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h-1))

	if big.NewInt(0).Add(vA, kA).Cmp(big.NewInt(0).Add(vb.Add(vb, kb), totalFee)) != 0 {
		panic(fmt.Sprintf("panic1 at hegiht %d", h))
	}
	totalIncome := big.NewInt(0)
	for i := int64(0); i < 3; i++ {
		res, err := instance.CurrentValidatorSet(&bind.CallOpts{BlockNumber: big.NewInt(h)}, big.NewInt(i))
		if err!=nil{
			fmt.Println(err)
		}
		totalIncome = totalIncome.Add(totalIncome, res.Incoming)
	}
	fmt.Println(totalIncome.String())
	fmt.Println(kA.String())

	end:=uint64(43)
	ite, err := instance.FilterValidatorDeposit(&bind.FilterOpts{Start: 43, End: &end}, nil)
	assert.NoError(t, err)
	for ite.Next() {
		fmt.Println(ite.Event.Validator.String())
	}

	if totalIncome.Cmp(kA) != 0 {
		panic(fmt.Sprintf("panic2 at hegiht %d", h))
	}
	fmt.Printf("check finish at height %d, txs %d, total income %s\n", h, len(txs), totalIncome.String())

}

func TestFee2(t *testing.T) {
	for h:=int64(0);h<100;h++{
		fmt.Printf("height %d\n",h)
		client, _ := ethclient.Dial(endpoint)
		instance, _ := NewValidator(validatorSetAddr, client)
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(h)))
		assert.NoError(t, err)
		txs := block.Transactions()
		totalFee := big.NewInt(0)
		for _, tx := range txs {
			price := tx.GasPrice()
			recip, err := client.TransactionReceipt(context.Background(), tx.Hash())
			assert.NoError(t, err)
			totalFee = totalFee.Add(totalFee, price.Mul(price, big.NewInt(int64(recip.GasUsed))))
		}
		vA, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h))
		kA, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h))
		vb, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h-1))
		kb, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h-1))

		if big.NewInt(0).Add(vA, kA).Cmp(big.NewInt(0).Add(vb.Add(vb, kb), totalFee)) != 0 {
			panic(fmt.Sprintf("panic1 at hegiht %d", h))
		}
		totalIncome := big.NewInt(0)
		for i := int64(0); i < 3; i++ {
			res, err := instance.CurrentValidatorSet(&bind.CallOpts{BlockNumber: big.NewInt(h)}, big.NewInt(i))
			fmt.Println(err)
			totalIncome = totalIncome.Add(totalIncome, res.Incoming)
		}
		if totalIncome.Cmp(kA) != 0 {
			fmt.Println(totalIncome.String())
			fmt.Println(kA.String())
			panic(fmt.Sprintf("panic2 at hegiht %d", h))
		}
	}
}


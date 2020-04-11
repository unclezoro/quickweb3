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
	headers:=make(chan  *types.Header,100)
	wsclient, err := ethclient.Dial(wsEndpoint)
	assert.NoError(t, err)
	wsclient.SubscribeNewHead(context.Background(),headers)
	for header:=range headers{
		h:= int64(header.Number.Uint64()-7)
		block,err:=client.BlockByNumber(context.Background(),big.NewInt(int64(h)))
		assert.NoError(t,err)
		txs:=block.Transactions()
		totalFee:=big.NewInt(0)
		for _,tx:=range txs{
			price:=tx.GasPrice()
			recip,err:=client.TransactionReceipt(context.Background(),tx.Hash())
			assert.NoError(t,err)
			totalFee = totalFee.Add(totalFee,price.Mul(price,big.NewInt(int64(recip.GasUsed))))
		}
		vA, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h))
		kA, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h))
		vb, _ := client.BalanceAt(context.Background(), systemRewardAddr, big.NewInt(h-1))
		kb, _ := client.BalanceAt(context.Background(), validatorSetAddr, big.NewInt(h-1))

		if big.NewInt(0).Add(vA,kA).Cmp(big.NewInt(0).Add(vb.Add(vb,kb),totalFee)) !=0{
			panic(fmt.Sprintf("panic1 at hegiht %d",h))
		}
		totalIncome:=big.NewInt(0)
		for i := int64(0); i < 3; i++ {
			res, _ := instance.CurrentValidatorSet(&bind.CallOpts{BlockNumber: big.NewInt(h)}, big.NewInt(i))
			totalIncome =totalIncome.Add(totalIncome,res.Incoming)
		}
		if totalIncome.Cmp(kA) !=0 {
			panic(fmt.Sprintf("panic2 at hegiht %d",h))
		}
		fmt.Printf("check finish at height %d, txs %d, total income %s\n",h,len(txs),totalIncome.String())
	}
}

func transferRoutin() {
	client, _ := ethclient.Dial(endpoint)
	for {
		sendEther(client, account, receiveAccount, big.NewInt(0).Mul(big.NewInt(params.GWei), big.NewInt(10000)), false)
		time.Sleep(1 * time.Second)
	}
}

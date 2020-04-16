package abi

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlashIndicate(t *testing.T){
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)
	slashIns,_:=NewSlash(slashAddr, client)
	for {
		_,count,_:=slashIns.GetSlashIndicator(nil, common.HexToAddress("0xa5f6a270f60c83624dD1849038eE7c9e8a3E55fc"))
		fmt.Println(count)
	}
}

func TestSlashIndicate1(t *testing.T){
	client, err := ethclient.Dial(endpoint)
	assert.NoError(t, err)
	slashIns,_:=NewSlash(slashAddr, client)

	x,_:=slashIns.MisdemeanorThreshold(nil)
	fmt.Println(x)

	govHub,_:=NewGov(govAddr, client)

	ite,_:=govHub.FilterFailReasonWithStr(&bind.FilterOpts{Start: 6689})
	if ite.Next(){
		fmt.Println(ite.Event.Message)
	}
}


func TestSystemReward(t *testing.T){
	client, _ := ethclient.Dial(endpoint)
	//assert.NoError(t, err)
	//rewardIns,_:=NewSystemReward(tokenHub, client)
	//ite,_:=rewardIns.FilterRewardTo(&bind.FilterOpts{Start: uint64(21533)},nil)
	//if ite.Next(){
	//	fmt.Println(ite.Event.To.String())
	//}
	//b,_:=client.BalanceAt(context.Background(),tokenHub,nil)
	//fmt.Println(b)
	r,_:=client.TransactionReceipt(context.Background(),common.HexToHash("0xae6a764bcbdb8311bb15a3a36a78f9ed8d4e069b134bc9b7b8bc6cd77e72c5ac"))
	for _,l:=range r.Logs{
		fmt.Println("address", l.Address.String())
		for _,topic:=range  l.Topics{
			fmt.Println("hash:", topic.String())
		}
	}
}
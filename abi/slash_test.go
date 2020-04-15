package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

var slashAddr = common.HexToAddress("0x0000000000000000000000000000000000001001")

func TestSlash(t *testing.T) {
	client, _ := ethclient.Dial(endpoint)
	instance, _ := NewSlash(slashAddr, client)

	validator,_:=NewValidator(validatorSetAddr,client)

	end:=uint64(76827)
	ite,err:=validator.FilterValidatorFelony(&bind.FilterOpts{Start: 76827,End: &end},nil,nil)
	if err!=nil{
		fmt.Println(err)
	}
	var slashAddr common.Address
	for ite.Next(){
		e:=ite.Event
		fmt.Println(e.Validator.String())
		slashAddr = ite.Event.Validator
	}

	_,count,_:=instance.GetSlashIndicator(&bind.CallOpts{BlockNumber: big.NewInt(76827)},slashAddr)
	fmt.Println(count.String())

	//
	total:=0
	for h:=uint64(76827);h>0;h--{
		ite,_:= instance.FilterValidatorSlashed(&bind.FilterOpts{Start: h,End: &h},[]common.Address{slashAddr})
		if ite.Next(){
			total++
			e:=ite.Event
			fmt.Printf("At height %v\n",e.Raw.BlockNumber)
		}
		if total == 150 {
			_,count,_:=instance.GetSlashIndicator(&bind.CallOpts{BlockNumber: big.NewInt(int64(h))},slashAddr)
			fmt.Println(count.String())
			break
		}
	}
}

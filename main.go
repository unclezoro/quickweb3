package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

func mulBig(v1 string, v2 string) *big.Int {
	x1 := math.MustParseBig256(v1)
	x2 := math.MustParseBig256(v2)
	return x1.Mul(x1, x2)
}

func fromWeiToEth(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), new(big.Float).SetInt(big.NewInt(params.Ether)))
}

func main() {
	fmt.Println(mulBig("0x4a817c800", "0x5208").String())

	fmt.Println(new(big.Int).Add(math.MustParseBig256("393750000000000"),math.MustParseBig256("0x17dfcdece400")))

	fmt.Println(math.MustParseBig256("0x989680").String())
}

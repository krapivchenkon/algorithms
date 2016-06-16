package main

import (
	"fmt"
	"math/big"
)

func main() {

	var N int64
	fmt.Scanf("%v\n", &N)
	Fact := big.NewInt(1)

	for i := int64(0); i < N; i++ {
		Fact.Mul(Fact, big.NewInt(i+1))
	}

	fmt.Println(Fact)

}

package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(int64(math.Pow10(15)))
	b := big.NewInt(int64(math.Pow10(10)))
	aMulb := big.NewInt(0).Mul(a, b)
	aDivb := big.NewInt(0).Div(a, b)
	aAddb := big.NewInt(0).Add(a, b)
	aSubb := big.NewInt(0).Sub(a, b)
	fmt.Println(a, b)
	fmt.Println("a * b =", aMulb)
	fmt.Println("a / b =", aDivb)
	fmt.Println("a + b =", aAddb)
	fmt.Println("a - b =", aSubb)

}

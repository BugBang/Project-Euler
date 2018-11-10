package main

import (
	"math/big"
	"strconv"
)

//2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
//What is the sum of the digits of the number 2**1000?
func main() {
	bigInt := new(big.Int)
	digits := bigInt.Exp(big.NewInt(2), big.NewInt(1000), nil)
	strDigits := digits.String()
	sum := 0
	for _, value := range strDigits {
		temp, _ := strconv.Atoi(string(value))
		sum += temp
	}
	println(sum)
}

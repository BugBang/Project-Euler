package main

import (
	"fmt"
	"math/big"
	"strconv"
)

//n! means n × (n − 1) × ... × 3 × 2 × 1
//
//For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
//and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
//Find the sum of the digits in the number 100!
func main() {
	var (
		mul    *big.Int = big.NewInt(int64(1))
		sum    int
		result int
		err    error
	)
	for i := 1; i <= 100; i++ {
		mul.Mul(mul, big.NewInt(int64(i)))
	}

	for _, value := range mul.String() {
		if sum, err = strconv.Atoi(string(value)); err == nil {
			result += sum
		}
	}

	fmt.Println(result)

}

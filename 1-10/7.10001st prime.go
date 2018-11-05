package main

import (
	"fmt"
	"math"
)

//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
//we can see that the 6th prime is 13.
//
//What is the 10 001st prime number?
func main() {
	var i int64 = 0
	count := 0
	for {
		if isPrime1(i) {
			count++
		}
		if count == 10001 {
			goto forend
		}
		i++
	}
forend:
	fmt.Println(i)
}

func isPrime1(a int64) bool {
	if a <= 1 {
		return false
	}

	for b := sqrt1(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a%b == 0 {
			return false
		}
	}

	return true
}

func sqrt1(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

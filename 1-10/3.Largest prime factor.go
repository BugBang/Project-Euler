package main

import (
	"fmt"
	"math"
)

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

func main() {
	var a int64 = 600851475143

	for b := sqrt(a); b >= 1; b-- {
		if a%b == 0 && isPrime(b) {
			fmt.Println(b)
			break
		}
	}
}

func isPrime(a int64) bool {
	if a <= 1 {
		return false
	}

	for b := sqrt(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a%b == 0 {
			return false
		}
	}

	return true
}

func sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

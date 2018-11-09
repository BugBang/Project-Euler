package main

import "fmt"

//The following iterative sequence is defined for the set of positive integers:
//
//n → n/2 (n is even)
//n → 3n + 1 (n is odd)
//
//Using the rule above and starting with 13, we generate the following sequence:
//
//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//
//Which starting number, under one million, produces the longest chain?
//
//NOTE: Once the chain starts the terms are allowed to go above one million.
func main() {
	var step int
	var value int
	for i := 0; i < 1000000; i++ {
		temp := chainLens(i)
		if temp > value {
			value = temp
			step = i
		}
	}

	fmt.Println(step)
}

func chainLens(value int) int {
	var step int = 1
	for value > 1 {
		if value%2 == 0 {
			value = value / 2
		} else {
			value = value*3 + 1
		}
		step++
	}
	return step
}

package main

import "fmt"

//2520 is the smallest number that can be divided
//by each of the numbers from 1 to 10 without any remainder.
//
//What is the smallest positive number that
//is evenly divisible by all of the numbers from 1 to 20?

func main() {
	x := 1
	for i := 1; i <= 20; i++ {
		if x%i != 0 {
			x++
			i = 1
		}
		if i == 20 {
			goto forend
		}
	}
forend:
	fmt.Println(x)
}

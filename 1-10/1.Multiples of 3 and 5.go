package main

import "fmt"

var sumX chan int

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
//
//Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
	//	sum(1000000000)
	//	sumX = make(chan int, 1)
	//	sums := 0
	//	count := 0
	//	for {
	//		select {
	//		case a := <-sumX:
	//			sums = sums + a
	//			count++
	//			if count == 2 {
	//				goto forend
	//			}
	//		}
	//	}
	//forend:
	//	fmt.Println(sums)
}

func sum(limit int) {
	go func() {
		sum3 := 0
		for i := 0; i < limit; i++ {
			if i%3 == 0 && i%5 != 0 {
				sum3 += i
			}
		}
		sumX <- sum3
	}()

	go func() {
		sum5 := 0
		for i := 0; i < limit; i++ {
			if i%5 == 0 {
				sum5 += i
			}
		}
		sumX <- sum5
	}()

}

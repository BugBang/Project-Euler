package main

import (
	"fmt"
	"math"
	"time"
)

//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//Find the sum of all the primes below two million.
func main() {
	sum := 0
	count := 0
	groups := 100
	var sumChan = make(chan int, 1)

	for i := 0; i < groups; i++ {
		go func(x int) {
			temp := 0
			for j := 2000000 / groups * x; j < 2000000/groups*(x+1); j++ {
				if isPrime2(int64(j)) {
					temp += j
				}
				if j == 2000000/groups*(x+1)-1 {
					sumChan <- temp
					sumChan <- -1
				}
			}
		}(i)

	}

	go func() {
		for {
			select {
			case a := <-sumChan:
				if a == -1 {
					count++
					if count == groups {
						goto forend
					}
					continue
				}
				sum += a
			}
		}
	forend:
		fmt.Println(sum)
	}()

	time.Sleep(time.Second * 1000)

}

func isPrime2(a int64) bool {
	if a <= 1 {
		return false
	}

	for b := sqrt2(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a%b == 0 {
			return false
		}
	}

	return true
}

func sqrt2(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

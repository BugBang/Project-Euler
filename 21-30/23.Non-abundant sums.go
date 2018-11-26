package main

import (
	"fmt"
	"math"
)

//并非盈数之和
//完全数是指真因数之和等于自身的那些数。
//例如，28的真因数之和为1 + 2 + 4 + 7 + 14 = 28，因此28是一个完全数。
//
//一个数n被称为亏数，如果它的真因数之和小于n；反之则被称为盈数。
//
//由于12是最小的盈数，它的真因数之和为1 + 2 + 3 + 4 + 6 = 16，
//所以最小的能够表示成两个盈数之和的数是24。通过数学分析可以得出，所有大于28123的数都可以被写成两个盈数的和；
//尽管我们知道最大的不能被写成两个盈数的和的数要小于这个值，但这是通过分析所能得到的最好上界。
//
//找出所有不能被写成两个盈数之和的正整数，并求它们的和。
func main() {
	ans := generateProfit()
	ansmap := map[int]int{}
	for _, a := range ans {
		ansmap[a] = 1
	}
	sum := 1
	for i := 2; i < 28124; i++ {

		found := false
		for _, a := range ans {
			if a > i {
				break
			}
			if ansmap[i-a] != 0 {
				found = true
				break
			}
		}

		if !found {
			sum += i
		}

	}

	fmt.Println(sum)
}

func generateProfit() (profits []int) {
	var (
		sqrt float64
		temp []int
		sum  int
	)
	for i := 1; i < 28124; i++ {
		temp = temp[0:0:0]
		sum = 0
		sqrt = math.Sqrt(float64(i))
		for j := 1; j <= int(sqrt); j++ {
			if i%j == 0 {
				temp = append(temp, j)
				if j != 1 && j != i/j {
					temp = append(temp, i/j)
				}
			}
		}
		for _, a := range temp {
			sum += a
		}

		if sum > i {
			profits = append(profits, i)
		}

	}
	return
}

package main

import (
	"strconv"
)

//可截素数
//
//3797有着奇特的性质。不仅它本身是一个素数，而且如果从左往右逐一截去数字，剩下的仍然都是素数：3797、797、97和7；
//同样地，如果从右往左逐一截去数字，剩下的也依然都是素数：3797、379、37和3。
//
//只有11个素数，无论从左往右还是从右往左逐一截去数字，剩下的仍然都是素数，求这些数的和。
//
//注意：2、3、5和7不被视为可截素数。

func main() {
	arr := make([]bool, 1000000)
	left := make([]bool, 1000000)
	right := make([]bool, 1000000)

	left[2], left[3], left[5], left[7] = true, true, true, true
	right[2], right[3], right[5], right[7] = true, true, true, true
	arr[0], arr[1], arr[6], arr[9] = true, true, true, true
	a := []int{3, 5, 7}
	for i := range a {
		for k := a[i] * 2; k < len(arr); k += a[i] {
			arr[k] = true
		}
	}

	var k, tmp int
	prime, counter, sum := 11, 0, 0
	for {
		if right[prime/10] {
			right[prime] = true
		}
		tmp, _ = strconv.Atoi(strconv.Itoa(prime)[1:])
		if left[tmp] {
			left[prime] = true
			if right[prime] {
				sum += prime
				counter++
				if counter == 11 {
					println(sum)
					return
				}
			}
		}
		for k = prime * 2; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
		} else {
			break
		}
	}
}

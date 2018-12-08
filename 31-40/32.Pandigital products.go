package main

import (
	"strconv"
	"strings"
)

//全数字的乘积
//
//如果一个n位数包含了1至n的所有数字恰好一次，我们称它为全数字的；例如，五位数15234就是1至5全数字的。
//
//7254是一个特殊的乘积，因为在等式39 × 186 = 7254中，被乘数、乘数和乘积恰好是1至9全数字的。
//
//找出所有被乘数、乘数和乘积恰好是1至9全数字的乘法等式，并求出这些等式中乘积的和。
//
//注意：有些乘积可能从多个乘法等式中得到，但在求和的时候只计算一次。

func main() {
	mp := map[int]int{}
	for i := 1; i < 200; i++ {
		for j := 1; j < 5000; j++ {
			if isPandigital(i, j, i*j) {
				mp[i*j] = 1
			}
		}
	}
	sum := 0
	for key, _ := range mp {
		sum += key
	}
	println(sum)
}

func isPandigital(mul1, mul2, result int) bool {
	str := strconv.Itoa(mul1) + strconv.Itoa(mul2) + strconv.Itoa(result)
	if len(str) != 9 {
		return false
	}
	for i := 1; i < 10; i++ {
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

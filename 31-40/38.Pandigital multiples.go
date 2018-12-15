package main

import (
	"strconv"
	"strings"
)

//全数字的倍数
//
//将192分别与1、2、3相乘：
//
//192 × 1 = 192
//192 × 2 = 384
//192 × 3 = 576
//
//连接这些乘积，我们得到一个1至9全数字的数192384576。我们称192384576为192和(1,2,3)的连接乘积。
//
//同样地，将9分别与1、2、3、4、5相乘，得到1至9全数字的数918273645，即是9和(1,2,3,4,5)的连接乘积。
//
//对于n > 1，所有某个整数和(1,2, … ,n)的连接乘积所构成的数中，最大的1至9全数字的数是多少？
func main() {
	max := 0
	var tmp int
	for i := 1; i < 10000; i++ {
		str := ""
		for j := 1; j < 9 && len(str) < 9; j++ {
			str += strconv.Itoa(i * j)
		}
		if len(str) == 9 && isPandigital2(str) {
			tmp, _ = strconv.Atoi(str)
			if tmp > max {
				max = tmp
			}
		}
	}
	println(max)
}

func isPandigital2(str string) bool {
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

package main

import "fmt"

//各位数字的阶乘
//
//145是个有趣的数，因为1! + 4! + 5! = 1 + 24 + 120 = 145。
//
//找出所有各位数字的阶乘和等于其本身的数，并求它们的和。
//
//注意：因为1! = 1和2! = 2不是和的形式，所以它们并不在讨论范围内。

var digit_fact = map[rune]int{}

func init() {
	digit_fact['0'] = 1
	digit_fact['1'] = 1
	digit_fact['2'] = 2
	fact := 2
	for i := 3; i < 10; i++ {
		digit_fact[rune(i+48)] = i * fact
		fact = i * fact
	}
}

func factsum(n int) int {
	str := fmt.Sprintf("%d", n)
	sum := 0
	for _, r := range str {
		sum += digit_fact[r]
	}
	return sum
}

func main() {
	sum := 0
	for i := 13; i < 7*362880; i++ {
		fs := factsum(i)
		if fs == i {
			sum += i
		}
	}
	fmt.Println(sum)
}

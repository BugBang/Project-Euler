package main

import (
	"fmt"
	"strconv"
)

//A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	result := palindromic()
	fmt.Printf("number1: %d\n", result[0])
	fmt.Printf("number2: %d\n", result[1])
	fmt.Printf("result : %d", result[2])
}

func palindromic() []int {
	var result []int = make([]int, 3)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if temp := i * j; isPalindromicNumber(temp) {
				if temp > result[2] {
					result[0] = i
					result[1] = j
					result[2] = temp
				}
			}
		}
	}
	return result
}

func isPalindromicNumber(number int) bool {
	s1 := strconv.Itoa(number)
	s2 := reverseString(s1)
	if s1 == s2 {
		return true
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

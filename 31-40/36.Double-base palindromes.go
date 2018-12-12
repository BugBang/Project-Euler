package main

import (
	"fmt"
	"strconv"
)

//双进制回文数
//
//十进制数585 = 10010010012（二进制表示），因此它在这两种进制下都是回文数。
//
//找出所有小于一百万，且在十进制和二进制下均回文的数，并求它们的和。
//
//（请注意，无论在哪种进制下，回文数均不考虑前导零。）
func main() {
	var (
		sum int
	)
	for i := 0; i <= 1000000; i++ {
		if isPalindromic(strconv.Itoa(i)) {
			if isPalindromic(Dec2Bin(i)) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}

func Dec2Bin(num int) (binString string) {
	binString = strconv.FormatInt(int64(num), 2)
	return
}

func isPalindromic(num string) bool {
	for i := 0; i < len(num); i++ {
		if num[i] != num[len(num)-i-1] {
			return false
		}
	}
	return true
}

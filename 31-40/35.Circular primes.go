package main

import (
	"fmt"
	"math"
	"strconv"
)

//圆周素数
//
//197被称为圆周素数，因为将它逐位旋转所得到的数：197/971和719都是素数。
//
//小于100的圆周素数有十三个：2、3、5、7、11、13、17、31、37、71、73、79和97。
//
//小于一百万的圆周素数有多少个？
func main() {
	var (
		count int
	)
	for i := 100; i <= 1000000; i++ {
		if isCircularPrime(i) {
			count++
		}
	}
	fmt.Println(count + 13)
}

func isCircularPrime(num int) bool {
	var (
		snum string
	)
	if !isPrime(num) {
		goto NOPRIME
	}
	snum = strconv.Itoa(num)
	snum = snum[:]
	for i := 1; i <= len(snum); i++ {
		s := snum[i:len(snum)] + snum[0:i]
		is, _ := strconv.Atoi(s)
		if !isPrime(is) {
			goto NOPRIME
		}
	}
	return true
NOPRIME:
	return false
}

func isPrime(in int) bool {
	limit := int(math.Sqrt(float64(in)))
	if in < 2 || in%2 == 0 {
		return false
	} else {
		for i := 2; i < limit; i++ {
			if in%i == 0 {
				return false
			}
		}
	}
	return true
}

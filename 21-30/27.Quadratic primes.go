package main

import "math"

//二次“素数生成”多项式
//
//欧拉发现了这个著名的二次多项式：
//
//n2 + n + 41
//对于连续的整数n从0到39，这个二次多项式生成了40个素数。然而，当n = 40时，402 + 40 + 41 = 40(40 + 1) + 41能够被41整除，同时显然当n = 41时，412 + 41 + 41也能被41整除。
//
//随后，另一个神奇的多项式n2 − 79n + 1601被发现了，对于连续的整数n从0到79，它生成了80个素数。这个多项式的系数-79和1601的乘积为-126479。
//
//考虑以下形式的二次多项式：
//
//n**2 + an + b, 满足|a| < 1000且|b| < 1000
//
//其中|n|指n的模或绝对值
//例如|11| = 11以及|−4| = 4
//
//这其中存在某个二次多项式能够对从0开始尽可能多的连续整数n都生成素数，求其系数a和b的乘积。
func main() {
	var maxCounter, aInMax, bInMax int
	maxCounter = 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
		inner:
			for n := 0; ; n++ {
				if !isPrime(n*n + a*n + b) {
					if n > maxCounter {
						maxCounter, aInMax, bInMax = n, a, b
					}
					break inner
				}
			}
		}
	}
	println(aInMax * bInMax)
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

package main

import (
	"fmt"
	"math/big"
)

//一千位斐波那契数
//斐波那契数列是按如下递归关系定义的数列：
//F1 = 1 F2 = 1
//Fn = Fn−1 + Fn−2
//
//因此斐波那契数列的前12项分别是：
//F1 = 1
//F2 = 1
//F3 = 2
//F4 = 3
//F5 = 5
//F6 = 8
//F7 = 13
//F8 = 21
//F9 = 34
//F10 = 55
//F11 = 89
//F12 = 144
//
//第一个有三位数字的项是第12项F12。
//
//在斐波那契数列中，第一个有1000位数字的是第几项？
func main() {
	var (
		fibNum *big.Int
		count  int
	)
	fn := fibonacci()
	for {
		fibNum = fn()
		count++
		if len(fibNum.String()) >= 1000 {
			goto forend
		}
	}
forend:
	fmt.Println(count - 1)

}

func fibonacci() func() *big.Int {
	a, b := big.NewInt(int64(-1)), big.NewInt(int64(1))
	return func() *big.Int {
		a, b = b, a.Add(a, b)
		return b
	}
}

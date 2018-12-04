package main

import "fmt"

//螺旋数阵对角线
//
//从1开始，按顺时针顺序向右铺开的5 × 5螺旋数阵如下所示：
//
//21 22 23 24 25
//20  7  8  9 10
//19  6  1  2 11
//18  5  4  3 12
//17 16 15 14 13
//可以验证，该数阵对角线上的数之和是101。
//
//以同样方式构成的1001 × 1001螺旋数阵对角线上的数之和是多少？
const LIMIT = 1001

func main() {
	i, s, w, sum := 1, 0, 1, 1
	for {
		for n := 0; n < 5; n++ {
			if n < 2 {
				s++
			}
			if n == 0 {
				w += 2
			} else {
				i += s
				sum += i
			}
		}
		if w == LIMIT {
			break
		}
	}
	fmt.Println(sum)
}

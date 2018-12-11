package main

import "fmt"

//消去数字的分数
//
//49/98是一个有趣的分数，因为缺乏经验的数学家可能在约简时错误地认为，等式49/98 = 4/8之所以成立，是因为在分数线上下同时抹除了9的缘故。
//
//我们也会想到，存在诸如30/50 = 3/5这样的平凡解。
//
//这类有趣的分数恰好有四个非平凡的例子，它们的分数值小于1，且分子和分母都是两位数。
//
//将这四个分数的乘积写成最简分数，求此时分母的值。

func main() {
	var i, j, k, ki, ij int
	result := 1.0
	for i = 1; i < 10; i++ {
		for j = 1; j < i; j++ {
			for k = 1; k < j; k++ {
				ki = 10*k + i
				ij = 10*i + j
				if ki*j == ij*k {
					result *= float64(ij) / float64(ki)
				}
			}
		}
	}
	fmt.Println(int(result))
}

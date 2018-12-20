package main

import "strconv"

//钱珀瑙恩常数
//
//将所有正整数连接起来构造的一个十进制无理数如下所示：
//
//0.123456789101112131415161718192021…
//可以看出小数点后第12位数字是1。
//
//如果dn表示上述无理数小数点后的第n位数字，求下式的值：
//
//d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

func main() {
	i, count, limit, mul := 1, 0, 10, 1
	var tmp int
	for limit < 1000001 {
		str := strconv.Itoa(i)
		count += len(str)
		if count >= limit {
			tmp, _ = strconv.Atoi(str[len(str)-(count-limit)-1 : len(str)-(count-limit)])
			mul *= tmp
			limit *= 10
		}
		i++
	}
	println(mul)
}

package main

import "fmt"

//倒数的循环节
//
//单位分数指分子为1的分数。分母为2至10的单位分数的十进制表示如下所示：
//
//1/2= 0.5
//1/3= 0.(3)
//1/4= 0.25
//1/5= 0.2
//1/6= 0.1(6)
//1/7= 0.(142857)
//1/8= 0.125
//1/9= 0.(1)
//1/10= 0.1
//
//这里0.1(6)表示0.166666…，括号内表示有一位循环节。可以看出，1/7有六位循环节。
//
//找出正整数d < 1000，其倒数的十进制表示小数部分有最长的循环节。
func main() {

	var (
		numerator int
		decimals  map[int]int
		zeroCount int
		temp      int
		result    int
	)

	for i := 2; i < 1000; i++ {
		zeroCount = 0
		numerator = 1
		decimals = make(map[int]int)
		for {
			if numerator < i {
				numerator *= 10
			} else {
				break
			}
		}
		if numerator%i == 0 {
			continue
		} else {
			numerator = 1
		}

		for {
			if numerator < i {
				numerator *= 10
				if numerator < i {
					zeroCount++
				}
			} else {
				break
			}
		}

		for {
			key := numerator - numerator/i*i
			if decimals[key]-1 == 0 {
				if len(decimals)+zeroCount > temp {
					temp = len(decimals) + zeroCount
					result = i
				}
				break
			}
			decimals[key] = 1
			numerator = key * 10
		}
	}
	fmt.Println(result)
}

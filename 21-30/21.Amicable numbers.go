//package main
//
//import "fmt"
//
////亲和数
////记d(n)为n的所有真因数（小于n且整除n的正整数）之和。
////如果d(a) = b且d(b) = a，且a ≠ b，那么a和b构成一个亲和数对，a和b被称为亲和数。
////
////例如，220的真因数包括1、2、4、5、10、11、20、22、44、55和110，因此d(220) = 284；
////而284的真因数包括1、2、4、71和142，因此d(284) = 220。
////
////求所有小于10000的亲和数的和。
//func main() {
//	var (
//		amicableNum1 int
//		amicableNum2 int
//		result       []int
//		sum          int
//	)
//	for i := 1; i < 100000; i++ {
//		if amicableNum1 = isAmicableNum(i); amicableNum1 != -1 {
//			if amicableNum2 = isAmicableNum(amicableNum1); amicableNum2 != -1 {
//				if amicableNum1 != amicableNum2 && amicableNum2 == i {
//					result = append(result, i)
//					result = append(result, amicableNum2)
//				}
//			}
//		}
//	}
//	result = removeDuplicateElement(result)
//	for _, v := range result {
//		sum += v
//	}
//	fmt.Println(sum)
//}
//
//func removeDuplicateElement(amicableNums []int) []int {
//	result := make([]int, 0, len(amicableNums))
//	temp := map[int]struct{}{}
//	for _, item := range amicableNums {
//		if _, ok := temp[item]; !ok {
//			temp[item] = struct{}{}
//			result = append(result, item)
//		}
//	}
//	return result
//}
//
//func getSum(a []int) (result int) {
//	for _, v := range a {
//		result += v
//	}
//	return
//}
//
//func isAmicableNum(a int) int {
//	var (
//		result []int
//	)
//	for i := 1; i < a; i++ {
//		if a%i == 0 {
//			result = append(result, i)
//		}
//	}
//	if len(result) > 0 {
//		return getSum(result)
//	}
//	return -1
//}

package main

import (
	"math"
)

func main() {
	sum, j := 0, 0
	for i := 1; i < 10000; i++ {
		j = sumOfProperDivisors(i)
		if i == sumOfProperDivisors(j) && i != j {
			sum += i
		}
	}
	println(sum)
}

func sumOfProperDivisors(input int) int {
	limit, sum := int(math.Sqrt(float64(input))), 0
	for i := 1; i <= limit; i++ {
		if input%i == 0 {
			if i == input/i {
				sum = sum + i
			} else {
				sum = sum + i + (input / i)
			}
		}
	}
	return sum - input
}

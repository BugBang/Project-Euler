package main

//各位数字的五次幂
//
//令人惊讶的是，只有三个数可以写成它们各位数字的四次幂之和：
//
//1634 = 1**4 + 6**4 + 3**4 + 4**4
//8208 = 8**4 + 2**4 + 0**4 + 8**4
//9474 = 9**4 + 4**4 + 7**4 + 4**4
//
//由于1 = 1**4不是一个和，所以这里并没有把它包括进去。
//
//这些数的和是1634 + 8208 + 9474 = 19316。
//
//找出所有可以写成它们各位数字的五次幂之和的数，并求这些数的和。
func main() {
	var (
		limit int
		sum   int
	)
	limit = 9 * 9 * 9 * 9 * 9 * 5
	for i := 11; i < limit; i++ {
		if i == sumOfFifthPowers(i) {
			sum += i
		}
	}
	println(sum)

}

func fifthPower(num int) int {
	return num * num * num * num * num
}

func sumOfFifthPowers(num int) int {
	sum := 0
	for num > 0 {
		sum += fifthPower(num % 10)
		num /= 10
	}
	return sum
}

package main

import (
	"fmt"
	"math/big"
)

//不同的幂
//
//考虑所有满足2 ≤ a ≤ 5和2 ≤ b ≤ 5的整数组合生成的幂a**b：
//
//2**2=4, 2**3=8, 2**4=16, 2**5=32
//3**2=9, 3**3=27, 3**4=81, 3**5=243
//4**2=16, 4**3=64, 4**4=256, 4**5=1024
//5**2=25, 5**3=125, 5**4=625, 5**5=3125
//
//如果把这些幂按照大小排列并去重，我们得到以下由15个不同的项组成的序列：
//
//4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125
//在所有满足2 ≤ a ≤ 100和2 ≤ b ≤ 100的整数组合生成的幂a**b排列并去重所得到的序列中，
//有多少个不同的项？
func main() {
	var (
		result    []string
		bigi      *big.Int
		bigj      *big.Int
		bigresult *big.Int
	)
	for i := 2; i <= 100; i++ {
		for j := 2; j <= 100; j++ {
			bigi = big.NewInt(int64(i))
			bigj = big.NewInt(int64(j))
			bigresult = new(big.Int).Exp(bigi, bigj, nil)
			result = append(result, bigresult.String())
		}
	}

	result = removeDuplicateElement(result)
	fmt.Println(len(result))
}

func removeDuplicateElement(data []string) []string {
	result := make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

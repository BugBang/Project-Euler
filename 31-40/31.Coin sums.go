package main

import "fmt"

/*硬币求和

英国的货币单位包括英镑£和便士p，在流通中的硬币一共有八种：

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), £2 (200p)

以下是组成£2的其中一种可行方式：

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

不限定使用的硬币数目，组成£2有多少种不同的方式？*/

var (
	sum   int
	coins []int
)

func main() {
	sum, coins = 200, []int{1, 2, 5, 10, 20, 50, 100, 200}
	fmt.Println(count(sum, len(coins)))
}

func count(n, m int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	} else if m <= 0 && n >= 1 {
		return 0
	}
	return count(n, m-1) + count(n-coins[m-1], m)
}

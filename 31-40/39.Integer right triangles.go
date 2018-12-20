package main

import "fmt"

//整数边长直角三角形
//
//若三边长{a,b,c}均为整数的直角三角形周长为p，当p = 120时，恰好存在三个不同的解：
//
//{20,48,52}, {24,45,51}, {30,40,50}
//在所有的p ≤ 1000中，p取何值时有解的数目最多？

func main() {
	answer := 0
	max := 0
	var solutions int
	for p := 5; p <= 1000; p++ {
		solutions = 0
		for c := p; c > p/3; c-- {
			for b := 1; b < c; b++ {
				a := p - c - b
				if isPythagoras(a, b, c) {
					solutions++
				}
			}
		}
		if solutions > max {
			max, answer = solutions, p
		}
	}
	fmt.Println(answer)
}

func isPythagoras(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"math"
)

//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
//a**2 + b**2 = c**2
//For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.
//
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.
func main() {
loop:
	for a := 1; a < 500; a++ {
		for b := a; b < 500; b++ {
			result, c := isSquare(a*a + b*b)
			if result && a+b+c == 1000 {
				fmt.Printf("a= %d\n", a)
				fmt.Printf("b= %d\n", b)
				fmt.Printf("x= %d\n", c)
				println(a * b * c)
				break loop
			}
		}
	}

	//for a := 1; a < 500; a++ {
	//	for b := a; b < 500; b++ {
	//		for c := b; c < 500; c++ {
	//			if a*a+b*b == c*c && a+b+c == 1000 {
	//				fmt.Printf("a= %d\n", a)
	//				fmt.Printf("b= %d\n", b)
	//				fmt.Printf("x= %d\n", c)
	//				println(a * b * c)
	//			}
	//		}
	//	}
	//}
}

func isSquare(num int) (bool, int) {
	sqrt := int(math.Sqrt(float64(num)))
	if sqrt*sqrt == num {
		return true, sqrt
	}
	return false, -1
}

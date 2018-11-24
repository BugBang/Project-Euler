package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

//姓名得分
//在这个46K的文本文件names.txt（右击并选择“目标另存为……”）中包含了五千多个姓名。
//首先将它们按照字母序排列，
//然后计算出每个姓名的字母值，乘以它在按字母顺序排列后的位置，以计算出姓名得分。
//
//例如，按照字母序排列后,
//位于第938位的姓名COLIN的字母值是3 + 15 + 12 + 9 + 14 = 53。
//因此，COLIN的姓名得分是938 × 53 = 49714。
//
//文件中所有姓名的姓名得分之和是多少？
func main() {

	var (
		score int
		sum   int
	)

	content, err := ioutil.ReadFile("code_kata/Project-Euler/21-30/22data")
	if err != nil {
		println("Error reading file")
		return
	}
	names := strings.Split(strings.Replace(string(content), "\"", "", -1), ",")
	sort.Strings(names)
	for i, v1 := range names {
		score = 0
		for _, v2 := range v1 {
			score += int(v2) - int('A') + 1
		}
		sum += score * (i + 1)
	}

	fmt.Println(sum)

}

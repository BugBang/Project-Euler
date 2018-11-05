package main

import "fmt"

//You are given the following information, but you may prefer to do some research for yourself.
//
//1 Jan 1900 was a Monday.
//Thirty days has September,
//April, June and November.
//All the rest have thirty-one,
//Saving February alone,
//Which has twenty-eight, rain or shine.
//And on leap years, twenty-nine.
//A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

var sundayCount int

func main() {
	var day int
	if isLeapYear(1900) {
		day = 4*30 + 7*31 + 29 + 1
	} else {
		day = 4*30 + 7*31 + 28 + 1
	}

	for year := 1901; year <= 2000; year++ {
		for month := 1; month < 12; month++ {
			if isLargeMonth(month) {
				day += 30
			} else if month != 2 {
				day += 31
			} else {
				if isLeapYear(year) {
					day += 29
				} else {
					day += 28
				}
			}
			if isSunday(day) {
				sundayCount++
			}
		}
		day += 31
		if isSunday(day) {
			sundayCount++
		}
	}

	fmt.Println(sundayCount)
}

func isLargeMonth(month int) bool {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return true
	}
	return false
}

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	} else if year%400 == 0 {
		return true
	}
	return false
}

func isSunday(day int) bool {
	if day%7 == 0 {
		return true
	}
	return false
}

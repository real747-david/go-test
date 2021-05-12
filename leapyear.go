package main

import "fmt"

var year int

func isLeapyear(num int) (result bool) {
	if num%100 == 0 {
		if num%400 == 0 {
			result = true
		} else {
			result = false
		}
	} else {
		if num%4 == 0 {
			result = true
		} else {
			result = false
		}
	}
	return result
}

func main() {
	fmt.Println("please input a number of year:")
	fmt.Scanln(&year)

	res := isLeapyear(year)
	if res {
		fmt.Printf("%d is leapyaer. \n", year)
	} else {
		fmt.Printf("%d is not leapyaer. \n", year)
	}
}

package main

import "fmt"

func x(a ...int) {
	var nums = []int { 3, 5, 7 }
	for _, v := range nums {
		correctValue := v - 1
		if len(a) > correctValue {
			fmt.Print(a[correctValue])
		} else {
			fmt.Print(0)
		}
	}
}

func main() {
	x(1, 2, 3, 4, 5)
}

package main

import (
	"fmt"
	"strconv"
)

func x(a ...int) string {
	var memoryValue string
	var nums = []int{3, 5, 7}
	for _, v := range nums {
		correctValue := v - 1
		if len(a) > correctValue {
			memoryValue += strconv.Itoa(a[correctValue])
		} else {
			memoryValue += "0"
		}
	}
	return memoryValue
}

func main() {
	fmt.Println(x(1, 2, 3, 4, 5))
}

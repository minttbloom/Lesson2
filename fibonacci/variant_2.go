package main

import "fmt"

func mainForVariant2() {
	fib(10)
}

func fib(count int) {
	array := []int { 0, 1 }
	for i := len(array); i < count; i++ {
		c := array[i-1] + array[i-2]
		array = append(array, c)
		fmt.Println(c)
	}
}
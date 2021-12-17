package main

import (
	"fmt"
)

type IntFunc func() int

func main() {
	counter := createFibonacciCounter()
	fmt.Println(counter(), counter(), counter(), counter(), counter(), counter(), counter())
}

func createFibonacciCounter() IntFunc {
	a := 0
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

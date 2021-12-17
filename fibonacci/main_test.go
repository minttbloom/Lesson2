package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFibonacciCounter(t *testing.T) {
	expectedResultArray := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	counter := createFibonacciCounter()
	for i := 0; i < 10; i++ {
		assert.Equal(t, expectedResultArray[i], counter())
	}
}

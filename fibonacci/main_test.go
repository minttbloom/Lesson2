package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFibonacciCounter(t *testing.T) {
	expectedResultArray := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	counter := createFibonacciCounter()
	notFibCounter := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		assert.Equal(t, expectedResultArray[i], counter())
	}
	assert.NotEqual(t, notFibCounter, counter())
}

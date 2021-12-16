package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestX(t *testing.T) {

	result1 := x(1, 6, 456, 7, 3, 6, 8, 3, 56, 7, 98, 23, 6, 7, 3, 1, 2, 3, 6, 8, 9, 4, 6, 8, 9, 6)
	result2 := x(1, 6, 6, 7, 3, 1, 2, 3, 6, 8, 9, 4, 6, 8, 9, 6)
	result3 := x(1, 6, 6, 7, 3, 1, 2)
	result4 := x(1, 6, 6, 7, 3, 1)
	result5 := x(1, 6, 6)
	result6 := x(1)

	assert.Equal(t, "45638", result1)
	assert.Equal(t, "632", result2)
	assert.Equal(t, "632", result3)
	assert.Equal(t, "630", result4)
	assert.Equal(t, "600", result5)
	assert.Equal(t, "000", result6)

}

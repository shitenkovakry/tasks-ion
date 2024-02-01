package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum_Average_Case1(t *testing.T) {
	array := []int{5, 2, 7, 6, 5, 3}

	expectedSum := 28
	expectedAverage := 4.67

	actualSum, actualAverage := FindSumAndAverageOfArray(array)

	assert.Equal(t, expectedSum, actualSum)
	assert.Equal(t, expectedAverage, actualAverage)
}

func Test_Sum_Average_Case2(t *testing.T) {
	array := []int{}

	expectedSum := 0
	expectedAverage := 0.00

	actualSum, actualAverage := FindSumAndAverageOfArray(array)

	assert.Equal(t, expectedSum, actualSum)
	assert.Equal(t, expectedAverage, actualAverage)
}

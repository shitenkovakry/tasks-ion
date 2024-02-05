package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Little_Combination_And_Average_Case1(t *testing.T) {
	array := []int{3, 1, 4, 5, 6, 7, 1, 2, 3}

	expected := []int{1, 2, 3}
	expectedLen := 2.00

	actual, actualLen := FindAverageOfLittleCOmbinationInArray(array)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedLen, actualLen)
}

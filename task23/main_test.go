package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Big_Combination_Case1(t *testing.T) {
	array := []int{3, 1, 4, 5, 6, 7, 1, 2, 3}

	expected := []int{4, 5, 6, 7}
	expectedLen := 4

	actual, actualLen := FindBigCombinationInArray(array)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedLen, actualLen)
}

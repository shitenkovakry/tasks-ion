package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Separate_Number_Case1(t *testing.T) {
	num := 93991

	expected := []int{9, 3, 9, 9, 1}

	actual := SeparateDigits(num)

	assert.Equal(t, expected, actual)
}

func Test_Add_Number_Case1(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	num := 93991

	expected := []int{9, 3, 1, 0, 2, 4, 5, 6, 7, 8}

	actual := AddNumbersToArray(array, num)

	assert.Equal(t, expected, actual)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Separate_Number_Case1(t *testing.T) {
	num := 321

	expected := []int{3, 2, 1}

	actual := SeparateDigits(num)

	assert.Equal(t, expected, actual)
}

func Test_Separate_Number_Case2(t *testing.T) {
	num := 3210

	expected := []int{3, 2, 1, 0}

	actual := SeparateDigits(num)

	assert.Equal(t, expected, actual)
}

func Test_Add_Number_Case1(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	num := 321

	expected := []int{3, 2, 1, 0, 4, 5, 6, 7, 8, 9}

	actual := AddNumbersToArray(array, num)

	assert.Equal(t, expected, actual)
}

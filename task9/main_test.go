package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Big_Odd_And_Even_Numbers_Case1(t *testing.T) {
	array := []int{3, 5, 1, 8, 9, 2, 7, 10, 13}

	expected := []int{10, 9}
	actual := FindBigEvenAndOddNumbers(array)

	assert.Equal(t, expected, actual)
}

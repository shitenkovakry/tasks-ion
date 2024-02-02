package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Small_Three_Odd_Numbers_Case1(t *testing.T) {
	array := []int{3, 5, 1, 8, 9, 2, 7, 10, 11, 0}

	expected := []int{1, 3, 5}
	actual := FindThreeSmallOddNumbers(array)

	assert.Equal(t, expected, actual)
}

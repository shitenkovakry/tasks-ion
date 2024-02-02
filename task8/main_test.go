package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Big_Three_Even_Numbers_Case1(t *testing.T) {
	array := []int{3, 4, 5, 8, 9, 1, 2, 7, 10, 0}

	expected := []int{10, 8, 4}
	actual := FindThreeBigEvenNumbers(array)

	assert.Equal(t, expected, actual)
}

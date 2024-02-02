package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Big_Three_Numbers_Case1(t *testing.T) {
	array := []int{3, 4, 5, 8, 9, 1, 2, 7, 10, 0}

	expected := []int{10, 9, 8}
	actual := FindThreeBigNumbers(array)

	assert.Equal(t, expected, actual)
}

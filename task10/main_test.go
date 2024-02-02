package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Summa_Between_Numbers_Case1(t *testing.T) {
	array := []int{2, -1, 6, 4, -3, 7, -2, 0, 9, 2}

	expected := 5
	actual := FindSumBetweenNumbers(array)

	assert.Equal(t, expected, actual)
}

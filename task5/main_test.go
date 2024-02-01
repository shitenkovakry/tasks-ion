package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reasult_Case1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}

	expectedOdd := []int{1, 3, 5}
	expectedEven := []int{2, 4}

	actualOdd, actualEven := DisplayEvenAndOddNumbers(array)

	assert.Equal(t, expectedOdd, actualOdd)
	assert.Equal(t, expectedEven, actualEven)
}

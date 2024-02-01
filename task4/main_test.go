package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Quantity_Positive_Case1(t *testing.T) {
	array := []int{1, -2, 4, -6, 7, -1, 3}

	expected := 4
	actual := CalculateQuantityOfPositiveValues(array)

	assert.Equal(t, expected, actual)
}

func Test_Quantity_Negative_Case1(t *testing.T) {
	array := []int{1, -2, 4, -6, 7, -1, 3}

	expected := 3
	actual := CalculateQuantityOfNegativeValues(array)

	assert.Equal(t, expected, actual)
}

func Test_Result_Positive_Case1(t *testing.T) {
	array := []int{1, -2, 4, -6, 7, -1, 3}

	expected := []int{4, 15}
	actual := ResultForPositiveValues(array)

	assert.Equal(t, expected, actual)
}

func Test_Result_Negative_Case1(t *testing.T) {
	array := []int{1, -2, 4, -6, 7, -1, 3}

	expected := []int{3, -9}
	actual := ResultForNegativeValues(array)

	assert.Equal(t, expected, actual)
}

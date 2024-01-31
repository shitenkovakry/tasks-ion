package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max_Case1(t *testing.T) {
	array := []int{3, 5, 8, 6, 2, 7, 10, 11}

	expected := 11
	actual := FindMaxInArray(array)

	assert.Equal(t, expected, actual)
}

func Test_Min_Case1(t *testing.T) {
	array := []int{3, 5, 8, 6, 2, 7, 10, 11}

	expected := 2
	actual := FindMinInArray(array)

	assert.Equal(t, expected, actual)
}

func Test_Reasult_Case1(t *testing.T) {
	array := []int{3, 5, 8, 6, 2, 7, 10, 11}

	expectedMax := "max value = 11\n"
	expectedMin := "min value = 2\n"

	actualMax, actualMin := GetResult(array)

	assert.Equal(t, expectedMax, actualMax)
	assert.Equal(t, expectedMin, actualMin)
}

func Test_Reasult_Case2(t *testing.T) {
	array := []int{}

	expectedMax := "array is empty"
	expectedMin := ""

	actualMax, actualMin := GetResult(array)

	assert.Equal(t, expectedMax, actualMax)
	assert.Equal(t, expectedMin, actualMin)
}

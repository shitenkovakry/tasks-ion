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

	expectedMax := 11
	expectedMin := 2

	actualMax, actualMin := GetResult(array)

	assert.Equal(t, expectedMax, actualMax)
	assert.Equal(t, expectedMin, actualMin)
}

func Test_Reasult_Case2(t *testing.T) {
	array := []int{}

	expectedMax := 0
	expectedMin := 0

	actualMax, actualMin := GetResult(array)

	assert.Equal(t, expectedMax, actualMax)
	assert.Equal(t, expectedMin, actualMin)
}

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

	expectedMin := 2
	expectedMax := 11

	actualMin, actualMax := GetResult(array)

	assert.Equal(t, expectedMin, actualMin)
	assert.Equal(t, expectedMax, actualMax)
}

func Test_Reasult_Case2(t *testing.T) {
	array := []int{}

	expectedMin := 0
	expectedMax := 0

	actualMin, actualMax := GetResult(array)

	assert.Equal(t, expectedMin, actualMin)
	assert.Equal(t, expectedMax, actualMax)
}

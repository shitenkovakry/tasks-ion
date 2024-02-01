package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Similar_Case1(t *testing.T) {
	array := []int{1, 3, 4, 3, 8}

	expected := 3
	actual := FindSimilarElements(array)

	assert.Equal(t, expected, actual)
}

func Test_Get_Index_Case1(t *testing.T) {
	array := []int{1, 3, 4, 3, 8}
	simValue := 3

	expected := []int{1, 3}
	actual := GetIndexesOfSimilarValues(array, simValue)

	assert.Equal(t, expected, actual)
}

func Test_Get_Result_Case1(t *testing.T) {
	array := []int{1, 3, 4, 3, 8}

	expected := []int{1, 3}
	actual := GetResult(array)

	assert.Equal(t, expected, actual)
}

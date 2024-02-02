package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Difference_Between_Numbers_Case1(t *testing.T) {
	array := []int{35, 78, 30, 200, 42}

	expected := 170
	actual := FindTheBigDifferenceBetweenTwoElements(array)

	assert.Equal(t, expected, actual)
}

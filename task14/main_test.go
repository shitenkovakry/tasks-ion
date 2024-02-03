package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case1(t *testing.T) {
	input := []int{45, 30, 45, 35, 67, 35, 30, 89, 90, 35}

	expected := Pairs{
		{
			3, 35,
		},
		{
			2, 30,
		},
		{
			2, 45,
		},
		{
			1, 67,
		},
		{
			1, 89,
		},
		{
			1, 90,
		},
	}

	actual := FindSimilarNumbers(input)
	assert.Equal(t, expected, actual)
}

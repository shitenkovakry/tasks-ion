package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse_Case1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}

	expected := []int{7, 6, 5, 4, 3, 2, 1}
	actual := ReverseArray(array)

	assert.Equal(t, expected, actual)
}

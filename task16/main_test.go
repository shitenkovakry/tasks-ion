package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Closer_To_Zero(t *testing.T) {
	array := []int{8, 3, 1, 9, -4, 7, 10, 3, 9}

	expected := []int{3, -4}

	actual := FindSumCloserToZero(array)

	assert.Equal(t, expected, actual)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Not_Intersecting_Numbers_Case1(t *testing.T) {
	array := []int{2, 3, 1, 7}
	arraySecond := []int{3, 5, 8, 2, 4, 6}

	expected := []int{1, 7}

	actual := FindNotIntersectingNumbers(array, arraySecond)

	assert.Equal(t, expected, actual)
}

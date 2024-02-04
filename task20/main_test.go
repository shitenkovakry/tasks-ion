package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_In_New_Array_Case1(t *testing.T) {
	array := []int{2, 3, 1, 7}
	arraySecond := []int{3, 5, 8, 2, 4, 6}

	expected := []int{2, 3, 1, 7, 5, 8, 4, 6}

	actual := AddFromArraiesToArray(array, arraySecond)

	assert.Equal(t, expected, actual)
}

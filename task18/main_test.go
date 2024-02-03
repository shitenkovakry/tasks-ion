package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_Number(t *testing.T) {
	array := []int{54, 2, 6, 7, 8, 65, 21, 91}
	numberShouldAdd := 23
	location := 4

	expected := []int{54, 2, 6, 7, 23, 8, 65, 21, 91}

	actual := AddNumberToArray(array, numberShouldAdd, location)

	assert.Equal(t, expected, actual)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Delete_Number(t *testing.T) {
	array := []int{35, 30, 89, 76, 155, 123}
	numberShouldDelete := 76

	expected := []int{35, 30, 89, 155, 123}

	actual := DeleteNumberFromArray(array, numberShouldDelete)

	assert.Equal(t, expected, actual)
}

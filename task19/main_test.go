package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Delete_Repetitive_Elements(t *testing.T) {
	array := []int{2, 3, 1, 2, 4, 3, 1, 2, 1, 3, 1, 3, 4}

	expected := []int{1, 2, 3, 4}

	actual := DeleteRepetitiveElementsFromArray(array)

	assert.Equal(t, expected, actual)
}

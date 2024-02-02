package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Result_Case1(t *testing.T) {
	array := []int{1, 2, 3, 2, 5, 2, 5}

	expectedSimValue := 2
	expectedQuantity := 3
	actualSimValue, actualQuantity := GetResult(array)

	assert.Equal(t, expectedSimValue, actualSimValue)
	assert.Equal(t, expectedQuantity, actualQuantity)
}

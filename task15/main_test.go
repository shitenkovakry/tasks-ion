package main

import (
	"testing"
)

func Test_case1(t *testing.T) {
	input := []int{4, 2, 3, 5, 1, 3, 8}
	number := 6

	actual := DetermineTwoElementsWhoseSumIsEqualToEteredNumber(input, number)
	t.Log(actual)
}

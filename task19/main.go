package main

import "sort"

func DeleteRepetitiveElementsFromArray(array []int) []int {
	result := []int{}
	lenOfArray := len(array)
	container := make(map[int]bool)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		container[value] = true
	}

	for value := range container {
		result = append(result, value)
	}

	sort.Slice(result, func(i, j int) bool {
		left := result[i]
		right := result[j]

		return left < right
	})
	return result
}

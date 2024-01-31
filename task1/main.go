package main

import "fmt"

func GetResult(array []int) (string, string) {
	if len(array) == 0 {
		return "array is empty", ""
	}

	maxValue := FindMaxInArray(array)
	minValue := FindMinInArray(array)

	maxValueString := fmt.Sprintln("max value =", maxValue)
	minValueString := fmt.Sprintln("min value =", minValue)

	return maxValueString, minValueString
}

func FindMaxInArray(array []int) int {
	max := array[0]

	for _, value := range array {
		if value > max {
			max = value
		}
	}

	return max
}

func FindMinInArray(array []int) int {
	min := array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
	}

	return min
}

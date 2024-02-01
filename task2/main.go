package main

import (
	"fmt"
	"strconv"
)

func FindSumAndAverageOfArray(array []int) (int, float64) {
	sum := 0
	average := 0.0
	lenOfArray := len(array)

	if lenOfArray == 0 {
		return 0, 0.0
	}

	for _, value := range array {
		sum += value
	}

	average = float64(sum) / float64(lenOfArray)

	averageStr := fmt.Sprintf("%.2f", average)

	average, _ = strconv.ParseFloat(averageStr, 64)

	return sum, average
}

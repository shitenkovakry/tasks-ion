package main

import (
	"fmt"
	"strconv"
)

func FindAverageOfLittleCOmbinationInArray(array []int) ([]int, float64) {
	arrayWithBigCombinations := [][]int{}
	arrayWithCombination := []int{array[0]}
	lenOfArray := len(array)
	summa := 0

	for index := 1; index < lenOfArray; index++ {
		value := array[index]

		if getLastElement(arrayWithCombination)+1 == value {
			arrayWithCombination = append(arrayWithCombination, value)
		} else {
			if len(arrayWithCombination) > 1 {
				copyofArrayWithCombination := make([]int, len(arrayWithCombination))
				copy(copyofArrayWithCombination, arrayWithCombination)

				arrayWithBigCombinations = append(arrayWithBigCombinations, copyofArrayWithCombination)
			}
			arrayWithCombination = []int{value}
		}
	}

	if len(arrayWithCombination) > 1 {
		arrayWithBigCombinations = append(arrayWithBigCombinations, arrayWithCombination)
	}

	minCombination := arrayWithBigCombinations[0]
	for _, combination := range arrayWithBigCombinations {
		if len(minCombination) > len(combination) {
			minCombination = combination
		}
	}

	for _, value := range minCombination {
		summa += value
	}

	average := float64(summa) / float64(len(minCombination))

	averageStr := fmt.Sprintf("%.2f", average)

	average, _ = strconv.ParseFloat(averageStr, 64)

	return minCombination, float64(average)
}

func getLastElement(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return array[len(array)-1]
}

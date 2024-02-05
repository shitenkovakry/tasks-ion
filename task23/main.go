package main

func FindBigCombinationInArray(array []int) ([]int, int) {
	arrayWithBigCombinations := [][]int{}
	arrayWithCombination := []int{array[0]}
	lenOfArray := len(array)

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

	maxCombination := []int{}
	for _, combination := range arrayWithBigCombinations {
		if len(maxCombination) < len(combination) {
			maxCombination = combination
		}
	}

	return maxCombination, len(maxCombination)
}

func getLastElement(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return array[len(array)-1]
}

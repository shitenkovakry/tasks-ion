package main

func FindSimilarElements(array []int) int {
	lenOfArray := len(array)

	for index := 0; index < lenOfArray; index++ {
		valueOfArray := array[index]
		for nextIndex := index + 1; nextIndex < lenOfArray; nextIndex++ {
			nextValueOfArray := array[nextIndex]

			if nextValueOfArray == valueOfArray {
				return valueOfArray
			}
		}
	}

	return 0
}

func GetIndexesOfSimilarValues(array []int, Simvalue int) []int {
	arrayOfIndexes := []int{}
	for index, value := range array {
		if value == Simvalue {
			arrayOfIndexes = append(arrayOfIndexes, index)
		}
	}

	return arrayOfIndexes
}

func GetResult(array []int) []int {
	simValue := FindSimilarElements(array)
	arrayOfIndexesWithSimilarValues := GetIndexesOfSimilarValues(array, simValue)

	return arrayOfIndexesWithSimilarValues
}

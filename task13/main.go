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

func GetIndexesOfSimilarValues(array []int, Simvalue int) int {
	arrayOfSimNumbers := []int{}
	for _, value := range array {
		if value == Simvalue {
			arrayOfSimNumbers = append(arrayOfSimNumbers, value)
		}
	}

	lenOfArrayOfSimilarNumbers := len(arrayOfSimNumbers)

	return lenOfArrayOfSimilarNumbers
}

func GetResult(array []int) (int, int) {
	simValue := FindSimilarElements(array)
	quantityOfSimValues := GetIndexesOfSimilarValues(array, simValue)

	return simValue, quantityOfSimValues
}

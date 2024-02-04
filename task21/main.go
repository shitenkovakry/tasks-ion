package main

func FindIntersectingNumbers(array []int, arraySecond []int) []int {
	arrayWithResult := []int{}
	lenOfSecondArray := len(arraySecond)

	for indexSecondArray := 0; indexSecondArray < lenOfSecondArray; indexSecondArray++ {
		valueSecondArray := arraySecond[indexSecondArray]

		isValueSecondInArrayResult := HasValueInArray(array, valueSecondArray)

		if isValueSecondInArrayResult {
			arrayWithResult = append(arrayWithResult, valueSecondArray)
		}
	}

	return arrayWithResult
}

func HasValueInArray(array []int, valueSecondArray int) bool {
	found := false

	for _, valueResult := range array {
		if valueResult == valueSecondArray {
			found = true

			break
		}
	}

	return found
}

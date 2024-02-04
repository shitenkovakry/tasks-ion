package main

func AddFromArraiesToArray(array []int, arraySecond []int) []int {
	arrayWithResult := []int{}
	lenOfSecondArray := len(arraySecond)

	for _, value := range array {
		arrayWithResult = append(arrayWithResult, value)
	}

	for indexSecondArray := 0; indexSecondArray < lenOfSecondArray; indexSecondArray++ {
		valueSecondArray := arraySecond[indexSecondArray]

		isValueSecondInArrayResult := HasValueInArray(arrayWithResult, valueSecondArray)

		if !isValueSecondInArrayResult {
			arrayWithResult = append(arrayWithResult, valueSecondArray)
		}
	}

	return arrayWithResult
}

func HasValueInArray(arrayWithResult []int, valueSecondArray int) bool {
	found := false

	for _, valueResult := range arrayWithResult {
		if valueResult == valueSecondArray {
			found = true

			break
		}
	}

	return found
}

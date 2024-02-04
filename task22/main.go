package main

func FindNotIntersectingNumbers(array []int, arraySecond []int) []int {
	arrayWithResult := []int{}
	lenOfArray := len(array)

	for indexArray := 0; indexArray < lenOfArray; indexArray++ {
		valueArray := array[indexArray]

		isValueInArrayResult := HasValueInArray(arraySecond, valueArray)

		if !isValueInArrayResult {
			arrayWithResult = append(arrayWithResult, valueArray)
		}
	}

	return arrayWithResult
}

func HasValueInArray(arraySecond []int, valueArray int) bool {
	found := false

	for _, valueResult := range arraySecond {
		if valueResult == valueArray {
			found = true

			break
		}
	}

	return found
}

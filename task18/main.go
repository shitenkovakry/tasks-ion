package main

func AddNumberToArray(array []int, numberShouldAdd int, location int) []int {
	arrayWithAddedNumber := []int{}
	lenOfArray := len(array)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]
		if index == location {
			arrayWithAddedNumber = append(arrayWithAddedNumber, numberShouldAdd, value)
		} else {
			arrayWithAddedNumber = append(arrayWithAddedNumber, value)
		}
	}

	return arrayWithAddedNumber
}

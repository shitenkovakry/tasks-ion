package main

func FindThreeBigNumbers(array []int) []int {
	arrayWithNumbers := []int{}
	maxNumber := array[0]
	maxNumberSecond := array[0]
	maxNumberThird := array[0]
	lenOfArray := len(array)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value > maxNumber {
			maxNumber = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, maxNumber)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value > maxNumberSecond && value < maxNumber {
			maxNumberSecond = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, maxNumberSecond)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value > maxNumberThird && value < maxNumberSecond {
			maxNumberThird = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, maxNumberThird)

	return arrayWithNumbers

}

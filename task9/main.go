package main

func FindBigEvenAndOddNumbers(array []int) []int {
	arrayWithNumbers := []int{}
	maxNumberEven := array[0]
	maxNumberOdd := array[0]
	lenOfArray := len(array)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 == 0 && value > maxNumberEven {
			maxNumberEven = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, maxNumberEven)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 != 0 && value > maxNumberOdd && value < maxNumberEven {
			maxNumberOdd = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, maxNumberOdd)

	return arrayWithNumbers
}

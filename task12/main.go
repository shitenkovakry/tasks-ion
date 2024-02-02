package main

func FindThreeSmallOddNumbers(array []int) []int {
	arrayWithNumbers := []int{}
	minNumber := array[0]
	lenOfArray := len(array)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 != 0 && value < minNumber {
			minNumber = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, minNumber)

	minNumberSecond := array[0]
	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 != 0 && value > minNumber {
			minNumberSecond = value

			break
		}
	}

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 != 0 && value < minNumberSecond && value > minNumber {
			minNumberSecond = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, minNumberSecond)

	minNumberThird := array[0]

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 != 0 && value > minNumberSecond {
			minNumberThird = value

			break
		}
	}

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		if value%2 != 0 && value < minNumberThird && value > minNumberSecond {
			minNumberThird = value
		}
	}

	arrayWithNumbers = append(arrayWithNumbers, minNumberThird)

	return arrayWithNumbers
}

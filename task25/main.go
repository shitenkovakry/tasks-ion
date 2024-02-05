package main

func AddNumbersToArray(array []int, num int) []int {
	arrayWithAddedNumbers := []int{}
	lenOfArray := len(array)

	separatedNumbers := SeparateDigits(num)

	for _, value := range separatedNumbers {
		arrayWithAddedNumbers = append(arrayWithAddedNumbers, value)
	}

	for indexArray := 0; indexArray < lenOfArray; indexArray++ {
		valueArray := array[indexArray]

		isValueInArrayWithAddedNumbers := HasValueInArray(arrayWithAddedNumbers, valueArray)

		if !isValueInArrayWithAddedNumbers {
			arrayWithAddedNumbers = append(arrayWithAddedNumbers, valueArray)
		}
	}

	return arrayWithAddedNumbers
}

func SeparateDigits(num int) []int {
	digits := make([]int, 0)

	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)

		num = num / 10
	}

	reverseDigits(digits)

	return digits
}

func reverseDigits(digits []int) {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
}

func HasValueInArray(array []int, value int) bool {
	found := false

	for _, valueResult := range array {
		if valueResult == value {
			found = true

			break
		}
	}

	return found
}

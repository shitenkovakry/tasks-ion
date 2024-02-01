package main

func CalculateQuantityOfPositiveValues(array []int) int {
	quantityOfPositive := 0
	for _, value := range array {
		if value > 0 {
			quantityOfPositive += 1
		}
	}

	return quantityOfPositive
}

func CalculateQuantityOfNegativeValues(array []int) int {
	quantityOfNegative := 0
	for _, value := range array {
		if value < 0 {
			quantityOfNegative += 1
		}
	}

	return quantityOfNegative
}

func ResultForPositiveValues(array []int) []int {
	arrayWithResult := []int{}
	summa := 0

	quantity := CalculateQuantityOfPositiveValues(array)

	for _, value := range array {
		if value > 0 {
			summa += value
		}
	}

	arrayWithResult = append(arrayWithResult, quantity)
	arrayWithResult = append(arrayWithResult, summa)

	return arrayWithResult
}

func ResultForNegativeValues(array []int) []int {
	arrayWithResult := []int{}
	summa := 0

	quantity := CalculateQuantityOfNegativeValues(array)

	for _, value := range array {
		if value < 0 {
			summa += value
		}
	}

	arrayWithResult = append(arrayWithResult, quantity)
	arrayWithResult = append(arrayWithResult, summa)

	return arrayWithResult
}

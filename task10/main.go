package main

const (
	from = -2
	to   = 4
)

func FindSumBetweenNumbers(array []int) int {
	arrayWithNecessaryNumbers := []int{}
	summa := 0

	for _, value := range array {
		if value >= from && value <= to {
			arrayWithNecessaryNumbers = append(arrayWithNecessaryNumbers, value)
		}
	}

	for _, value := range arrayWithNecessaryNumbers {
		summa += value
	}

	return summa
}

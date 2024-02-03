package main

func DeleteNumberFromArray(array []int, numberShouldDelete int) []int {
	arrayWithoutDeletedNumber := []int{}

	for _, value := range array {
		if value == numberShouldDelete {
			continue
		}

		arrayWithoutDeletedNumber = append(arrayWithoutDeletedNumber, value)
	}

	return arrayWithoutDeletedNumber
}

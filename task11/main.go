package main

func FindTheBigDifferenceBetweenTwoElements(array []int) int {
	max := array[0]
	min := array[0]

	for _, value := range array {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	difference := max - min

	return difference
}

package main

func GetResult(array []int) (int, int) {
	if len(array) == 0 {
		return 0, 0
	}

	minValue := FindMinInArray(array)
	maxValue := FindMaxInArray(array)

	return minValue, maxValue
}

func FindMaxInArray(array []int) int {
	max := array[0]

	for _, value := range array {
		if value > max {
			max = value
		}
	}

	return max
}

func FindMinInArray(array []int) int {
	min := array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
	}

	return min
}

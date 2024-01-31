package main

func FindMinMaxInArray(array []int) (int, int) {
	if len(array) == 0 {
		return 0, 0
	}

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

	return min, max
}

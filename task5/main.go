package main

func DisplayEvenAndOddNumbers(array []int) ([]int, []int) {
	arrayWithEvens := []int{}
	arrayWithOdd := []int{}

	for _, value := range array {
		if value%2 == 0 {
			arrayWithEvens = append(arrayWithEvens, value)
		} else {
			arrayWithOdd = append(arrayWithOdd, value)
		}
	}

	return arrayWithOdd, arrayWithEvens

}

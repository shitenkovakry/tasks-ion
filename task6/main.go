package main

func ReverseArray(array []int) []int {
	reversedArray := []int{}
	lenOfArray := len(array)

	for index := lenOfArray - 1; index <= lenOfArray; index-- {
		if index < 0 {
			break
		}

		element := array[index]
		reversedArray = append(reversedArray, element)
	}

	return reversedArray
}

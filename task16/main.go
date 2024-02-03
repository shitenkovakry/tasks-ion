package main

func FindSumCloserToZero(array []int) []int {
	arrayWithAbsoluteSumma := []int{}
	lenOfArray := len(array)
	closestToZero := array[0] + array[1]
	resultA := 0
	resultB := 0

	for index := 0; index < lenOfArray; index++ {
		for nextIndex := index + 1; nextIndex < lenOfArray; nextIndex++ {

			summa := array[index] + array[nextIndex]

			absSum := summa
			if summa < 0 {
				absSum = -summa
			}

			closestToZeroAbs := closestToZero
			if closestToZero < 0 {
				closestToZeroAbs = -closestToZero
			}

			if absSum < closestToZeroAbs {
				closestToZero = summa
				resultA, resultB = array[index], array[nextIndex]
			}
		}
	}

	arrayWithAbsoluteSumma = append(arrayWithAbsoluteSumma, resultA, resultB)

	return arrayWithAbsoluteSumma
}

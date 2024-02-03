package main

type Result struct {
	Values  []int
	Indexes []int
}

type Results []*Result

func DetermineTwoElementsWhoseSumIsEqualToEteredNumber(array []int, equalNumber int) Results {
	lenOfArray := len(array)
	results := make(Results, 0)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		for nextIndex := index + 1; nextIndex < lenOfArray; nextIndex++ {
			nextValue := array[nextIndex]

			sumValues := value + nextValue

			if sumValues == equalNumber {

				results = append(results, &Result{
					Values:  []int{value, nextValue},
					Indexes: []int{index, nextIndex},
				})
			}
		}
	}

	return results
}

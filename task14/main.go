package main

import (
	"sort"
)

type Pair struct {
	NumberOfOccurences int
	Value              int
}

type Pairs []*Pair

func FindSimilarNumbers(array []int) Pairs {
	lenOfArray := len(array)
	result := make(map[int]int)

	for index := 0; index < lenOfArray; index++ {
		value := array[index]

		_, hasValue := result[value]
		if hasValue {
			result[value]++
		} else {
			result[value] = 1
		}

	}

	pairs := make(Pairs, 0)

	for value, numberOfOccurences := range result {
		pairs = append(pairs, &Pair{
			Value:              value,
			NumberOfOccurences: numberOfOccurences,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		left := pairs[i]
		right := pairs[j]

		if left.NumberOfOccurences > right.NumberOfOccurences {
			return true
		}

		if left.NumberOfOccurences < right.NumberOfOccurences {
			return false
		}

		return left.Value < right.Value
	})

	return pairs
}

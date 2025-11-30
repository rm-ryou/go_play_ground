package problem

import "math"

func CountValue(a []int, key int) int {
	count := 0

	for _, v := range a {
		if v == key {
			count++
		}
	}

	return count
}

func FetchSecondMinValue(a []int) int {
	firstMinValue := math.MaxInt
	secondMinValue := math.MaxInt

	for _, v := range a {
		if v < firstMinValue {
			secondMinValue = firstMinValue
			firstMinValue = v
		} else if v < secondMinValue {
			secondMinValue = v
		}
	}

	return secondMinValue
}

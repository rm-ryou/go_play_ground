package main

import "math"

func BasicLinearSearch(ary []int, key int) bool {
	flg := false

	for _, v := range ary {
		if v == key {
			flg = true
		}
	}

	return flg
}

func FindIndex(ary []int, key int) int {
	idx := -1

	for i, v := range ary {
		if v == key {
			idx = i
		}
	}

	return idx
}

func FindMinValue(ary []int) int {
	minValue := math.MaxInt

	for _, v := range ary {
		if minValue > v {
			minValue = v
		}
	}

	return minValue
}

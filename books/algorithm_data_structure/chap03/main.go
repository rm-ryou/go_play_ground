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

func FindMinPairAboveNum(a, b []int, k int) int {
	minValue := math.MaxInt

	for _, va := range a {
		for _, vb := range b {
			if va+vb > k && va+vb < minValue {
				minValue = va + vb
			}
		}
	}
	return minValue
}

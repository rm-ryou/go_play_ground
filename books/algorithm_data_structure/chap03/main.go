package main

import (
	"math"
)

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

func IsCreateValueFromAry(a []int, w int) bool {
	flg := false
	// w = 10
	// a = [1, 2, 4, 5]
	// ans = true
	// len(a) = 4
	// 0000
	// 0001
	// 0010
	// 0011
	// 0100
	// ...

	n := len(a)
	for bit := range 1<<n - 1 {
		sum := 0
		for j := range n {
			if bit&(1<<j) != 0 {

				sum += a[j]
			}
		}
		if sum == w {
			flg = true
		}
	}
	return flg
}

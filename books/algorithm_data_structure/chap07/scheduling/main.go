package main

import "slices"

const (
	START int = iota
	END
)

func scheduling(arry [][]int) int {
	slices.SortFunc(arry, func(a, b []int) int {
		if a[END] < b[END] {
			return 1
		} else {
			return -1
		}
	})

	res := 0
	currentEndTime := 0
	for _, v := range arry {
		if v[START] < currentEndTime {
			continue
		}

		res++
		currentEndTime = v[END]
	}
	return res
}

func main() {
}

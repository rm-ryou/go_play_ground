package main

import (
	"slices"
	"sort"
)

const (
	WORK int = iota
	TIME
)

func megalo(n int, d, t []int) bool {
	arry2d := make([][]int, n)
	for i := range n {
		arry2d[i] = make([]int, 2)
		arry2d[i][0] = d[i]
		arry2d[i][1] = t[i]
	}

	slices.SortFunc(arry2d, func(a, b []int) int {
		if a[TIME] > b[TIME] {
			return 1
		} else {
			return -1
		}
	})

	res := true
	currentEndTime := 0
	for _, v := range arry2d {
		if currentEndTime+v[WORK] > v[TIME] {
			res = false
		}

		currentEndTime += v[WORK]
	}

	return res
}

func plan2d(n int, red, blue [][]int) int {
	isUsed := make([]bool, n)
	slices.SortFunc(blue, func(a, b []int) int {
		if a[1] > b[1] {
			return 1
		} else {
			return -1
		}
	})

	res := 0
	for i, v := range blue {
		maxY, maxId := -1, -1
		for j, vRed := range red {
			if isUsed[j] {
				continue
			}

			if (vRed[0] >= v[0]) || (vRed[1] >= v[1]) {
				continue
			}

			if maxY < vRed[1] {
				maxY = vRed[1]
				maxId = j
			}
		}

		if maxId == -1 {
			continue
		}

		res++
		isUsed[i] = true
	}
	return res
}

func prob71(n int, a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	res := 0
	for _, v := range b {
		if a[res] < v {
			res += 1
		}
	}
	return res
}

func main() {}

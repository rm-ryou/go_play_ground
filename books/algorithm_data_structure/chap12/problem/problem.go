package problem

import (
	"slices"
)

func Energy(n, m int, price, count []int) int {
	info := make([][]int, n)
	for i := range info {
		info[i] = make([]int, 2)
		info[i][0] = price[i]
		info[i][1] = count[i]
	}

	slices.SortFunc(info, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else {
			return +1
		}
	})

	var res int
	for _, v := range info {
		if m >= v[1] {
			res += v[0] * v[1]
			m -= v[1]
		} else {
			res += v[0] * m
			break
		}
	}

	return res
}

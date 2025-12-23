package problem

import (
	"sort"
)

func Pair(a, b []int) int {
	var count int

	sort.Ints(a)
	sort.Ints(b)

	for _, v := range b {
		if a[count] < v {
			count++
		}
	}

	return count
}

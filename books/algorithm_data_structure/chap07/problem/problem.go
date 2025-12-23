package problem

import "sort"

func Pair(a, b []int) int {
	var count int

	sort.Ints(a)
	sort.Ints(b)

	idxB := 0
	for i := range a {
		if idxB > len(b) {
			break
		}

		j := idxB
		for ; j < len(b); j++ {
			if a[i] < b[j] {
				count++
			}
		}
		idxB = j
	}

	return count
}

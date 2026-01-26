package problem

import (
	"fmt"
	"slices"
	"sort"
)

func ComplessMap(n int, a []int) []int {
	cAry := make([]int, n)
	copy(cAry, a)

	slices.SortFunc(cAry, func(a, b int) int {
		if a > b {
			return 1
		} else {
			return -1
		}
	})

	res := make([]int, n)
	for i := range a {
		idx := sort.Search(len(a), func(j int) bool {
			return cAry[j] >= a[i]
		})

		res[i] = idx
	}

	return res
}

// ai, < bj < ckを満たすi,j,kの組み合わせが何個あるか
func Sunke(n int, a, b, c []int) int {
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var res int
	for _, v := range b {
		// b未満のaの個数
		underVIdx := sort.Search(n, func(x int) bool {
			return a[x] >= v
		})

		upperVIdx := n - sort.Search(n, func(x int) bool {
			return c[x] > v
		})

		res += underVIdx * upperVIdx
	}
	return res
}

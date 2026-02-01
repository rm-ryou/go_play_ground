package problem

import (
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

func Darts(n, m int, a []int) int {
	a = append(a, 0)
	sumOfA := make([]int, 0)
	for i := range a {
		for j := range a {
			sumOfA = append(sumOfA, a[i]+a[j])
		}
	}

	sort.Ints(sumOfA)
	var res int
	for i := range sumOfA {
		idx := sort.Search(len(sumOfA), func(x int) bool {
			return sumOfA[x] >= m-sumOfA[i]
		})
		if idx == len(sumOfA) || idx == 0 {
			continue
		}

		if res <= sumOfA[i]+sumOfA[idx-1] {
			res = sumOfA[i] + sumOfA[idx-1]
		}
	}

	return res
}

func AggressiveCows(n, m int, a []int) int {
	sort.Ints(a)

	left, right := 0, int(1<<10)
	for right-left > 1 {
		mid := (right + left) / 2

		cnt := 1
		prevIdx := 0
		for i := range n {
			if a[i]-a[prevIdx] >= mid {
				cnt++
				prevIdx = i
			}
		}

		if cnt >= m {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

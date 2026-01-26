package main

import (
	"math"
	"slices"
	"sort"
)

func CalcMostMinPair(n, k int, a, b []int) int {
	var res int = 10e6
	cAry := make([]int, len(b))
	copy(cAry, b)

	slices.SortFunc(cAry, func(a, b int) int {
		if a > b {
			return 1
		} else {
			return -1
		}
	})

	for _, v := range a {
		if v > k {
			continue
		}
		tmpIdx := sort.Search(len(cAry), func(i int) bool {
			return cAry[i] >= k-v
		})
		tmp := v + cAry[tmpIdx]
		if tmp < res {
			res = tmp
		}
	}
	return res
}

func Shooting(n int, h, s []int) int {
	l, r := 0, math.MaxInt

	for r-l > 1 {
		mid := (r-l)/2 + l
		// 高さがmid以内に破り切れるか
		canShooting := true
		value := make([]int, n)
		for i := range n {
			if mid < h[i] {
				canShooting = false
			} else {
				// mid以内に破るまでの制限時間
				value[i] = (mid - h[i]) / s[i]
			}
		}

		slices.SortFunc(value, func(a, b int) int {
			if a > b {
				return 1
			} else {
				return -1
			}
		})

		for i := range n {
			if value[i] < i {
				canShooting = false
			}
		}

		if canShooting {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

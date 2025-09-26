package main

import (
	"math"
	"sort"
)

func sniper(n int, h, s []int) int {
	left, right := 0, math.MaxInt
	for right-left > 1 {
		mid := (left + right) / 2

		flg := true
		t := make([]int, n)
		for i := range n {
			if mid < h[i] {
				flg = false
			} else {
				t[i] = (mid - h[i]) / s[i]
			}
		}
		sort.Ints(t)
		for i := range n {
			if t[i] < i {
				flg = false
			}
		}
		if flg {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func main() {}

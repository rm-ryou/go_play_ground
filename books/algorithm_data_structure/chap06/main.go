package main

import (
	"fmt"
	"math"
	"sort"
)

func BinarySearch(ary []int, key int) bool {
	left, right := 0, len(ary)-1

	for right >= left {
		mid := left + (right-left)/2
		if ary[mid] == key {
			return true
		} else if ary[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func IsGuessNumberUnderN(min, max, limit int) bool {
	var turn int

	for max > min {
		// 少数以下は切り捨てるため、最悪ケースは常に右を撮り続けること
		mid := min + (max-min)/2
		min = mid + 1
		turn++
	}

	return turn <= limit
}

func lowerBound(ary []int, key int) int {
	left, right := 0, len(ary)-1

	for right-left > 1 {
		mid := left + (right-left)/2
		if ary[mid] >= key {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

func MinSumOverK(a, b []int, k int) int {
	res := math.MaxInt

	sort.Ints(a)
	sort.Ints(b)

	for _, v := range a {
		if v > k {
			continue
		}
		tmp := v + b[lowerBound(b, k-v)]
		fmt.Println(tmp)

		if res > tmp {
			res = tmp
		}
	}

	return res
}

func Shooting(n int, height, speed []int) int {
	left, right := 0, math.MaxInt

	for right-left > 1 {
		canShooting := true
		mid := (left + right) / 2

		timeLimit := make([]int, n)
		for i := range timeLimit {
			if mid < height[i] {
				canShooting = false
			} else {
				timeLimit[i] = (mid - height[i]) / speed[i]
			}
		}

		sort.Ints(timeLimit)
		for i, v := range timeLimit {
			if v < i {
				canShooting = false
			}
		}

		if canShooting {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

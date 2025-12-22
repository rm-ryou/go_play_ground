package main

import (
	"math"
	"slices"
)

func Coin(value int, nums []int) int {
	var res int

	mapNums := []int{50, 10, 5, 1}

	for i, v := range mapNums {
		canUse := value / v
		canUse = int(math.Min(float64(canUse), float64(nums[i])))

		res += canUse
		value -= canUse * v
	}

	return res
}

func Schedule(N int, starts, ends []int) int {
	dAry := make([][]int, N)
	for i := range dAry {
		dAry[i] = make([]int, 2)
		dAry[i][0] = starts[i]
		dAry[i][1] = ends[i]
	}

	slices.SortFunc(dAry, func(a, b []int) int {
		if a[1] > b[1] {
			return 1
		} else {
			return -1
		}
	})

	res := 0
	currentEndTime := 0
	for i := range dAry {
		if dAry[i][0] < currentEndTime {
			continue
		}

		currentEndTime = dAry[i][1]
		res++
	}
	return res
}

func MultipleArray(n int, buttons, nums []int) int {
	res := 0
	sum := 0
	for i := n - 1; i >= 0; i-- {
		buttons[i] += sum
		d := buttons[i] % nums[i]

		currentCount := 0
		if d != 0 {
			currentCount = nums[i] - d
		}

		res += currentCount
		sum = res
	}

	return res
}

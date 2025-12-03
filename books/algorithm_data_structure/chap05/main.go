package main

import "math"

func Flog(n int, ary []int) int {
	dp := make([]int, n)
	for i := range dp {
		if i != 0 {
			dp[i] = math.MaxInt
		}
	}

	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = int(math.Abs(float64(ary[i] - ary[i-1])))
		} else {
			dp[i] = int(
				math.Min(
					float64(dp[i-1]+int(math.Abs(float64(ary[i]-ary[i-1])))),
					float64(dp[i-2]+int(math.Abs(float64(ary[i]-ary[i-2])))),
				),
			)
		}
	}

	return dp[n-1]
}

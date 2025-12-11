package main

import (
	"math"
)

func chMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func Flog1(N int, height []int) int {
	dp := make([]int, N+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	for i := range N {
		if i == 1 {
			dp[i] = abs(height[i] - height[i-1])
		}
		if i > 1 {
			dp[i] = chMin(
				dp[i-1]+abs(height[i]-height[i-1]),
				dp[i-2]+abs(height[i]-height[i-2]),
			)
		}
	}

	return dp[N-1]
}

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
	dp := make([]int, N)
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

func Flog2(N, K int, height []int) int {
	dp := make([]int, N)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	for i := 1; i < N; i++ {
		for j := 1; j <= K; j++ {
			if i == 1 {
				dp[i] = abs(height[i] - height[i-1])
			} else if i >= j {
				dp[i] = chMin(dp[i], dp[i-j]+abs(height[i]-height[i-j]))
			}
		}
	}

	return dp[N-1]
}

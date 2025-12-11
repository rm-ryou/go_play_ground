package main

import (
	"math"
)

func chMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func chMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
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

func Vacation(N int, a, b, c []int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, 3)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	for i := range N + 1 {
		if i >= 1 {
			dp[i][0] = chMax(dp[i-1][1], dp[i-1][2]) + a[i-1]
			dp[i][1] = chMax(dp[i-1][0], dp[i-1][2]) + b[i-1]
			dp[i][2] = chMax(dp[i-1][0], dp[i-1][1]) + c[i-1]
		}
	}

	return chMax(dp[N][0], chMax(dp[N][1], dp[N][2]))
}

func Knapsack1(N, W int, w, v []int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := range N {
		for j := range W + 1 {
			if j >= w[i] {
				dp[i+1][j] = chMax(dp[i+1][j], dp[i][j-w[i]]+v[i])
			}
			dp[i+1][j] = chMax(dp[i+1][j], dp[i][j])
		}
	}

	return dp[N][W]
}

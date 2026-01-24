package main

import (
	"fmt"
	"math"
)

func Flog(N int, h []int) int {
	dp := make([]int, 7)
	for i := range dp {
		dp[i] = 10e6
	}

	dp[0] = 0
	for i := 1; i < N; i++ {
		if i > 1 {
			dp[i] = int(
				math.Min(
					float64(dp[i-1])+math.Abs(float64(h[i]-h[i-1])),
					float64(dp[i-2])+math.Abs(float64(h[i]-h[i-2])),
				),
			)
		}
	}

	return dp[N-1]
}

func FlogPushBased(N int, h []int) int {
	const inf = int(1 << 60)
	dp := make([]int, N)
	for i := range dp {
		dp[i] = inf
	}

	dp[0] = 0
	for i := range h {
		if i+1 < N {
			dp[i+1] = int(
				math.Min(
					float64(dp[i+1]),
					float64(dp[i])+math.Abs(float64(h[i]-h[i+1])),
				),
			)
		}
		if i+2 < N {
			dp[i+2] = int(
				math.Min(
					float64(dp[i+2]),
					float64(dp[i])+math.Abs(float64(h[i]-h[i+2])),
				),
			)
		}
	}

	return dp[N-1]
}

func Knapsack(N int, w []int, v []int, W int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := range N {
		for j := range W + 1 {
			if j-w[i] >= 0 {
				dp[i+1][j] = int(
					math.Max(
						float64(dp[i+1][j]),
						float64(dp[i][j-w[i]]+v[i]),
					),
				)
			}

			dp[i+1][j] = int(
				math.Max(
					float64(dp[i+1][j]),
					float64(dp[i][j]),
				),
			)
		}
	}

	return dp[N][W]
}

func EditDistance(S, T string) int {
	const inf = int(1 << 60)
	dp := make([][]int, len(S)+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0
	for i := range len(S) + 1 {
		for j := range len(T) + 1 {
			if i > 0 && j > 0 {
				if S[i-1] == T[j-1] {
					dp[i][j] = int(
						math.Min(
							float64(dp[i][j]),
							float64(dp[i-1][j-1]),
						),
					)
				} else {
					dp[i][j] = int(
						math.Min(
							float64(dp[i][j]),
							float64(dp[i-1][j-1]+1),
						),
					)
				}
			}

			if i > 0 {
				dp[i][j] = int(
					math.Min(
						float64(dp[i][j]),
						float64(dp[i-1][j]+1),
					),
				)
			}

			if j > 0 {
				dp[i][j] = int(
					math.Min(
						float64(dp[i][j]),
						float64(dp[i][j-1]+1),
					),
				)
			}
		}
	}

	fmt.Println(dp)
	return dp[len(S)][len(T)]
}

package main

import (
	"math"
)

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

func chMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func Knapsack(N, W int, wAry, vAry []int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := range N {
		for w := range W + 1 {
			if w-wAry[i] >= 0 {
				dp[i+1][w] = chMax(dp[i+1][w], dp[i][w-wAry[i]]+vAry[i])
			}
			dp[i+1][w] = chMax(dp[i+1][w], dp[i][w])
		}
	}

	return dp[N][W]
}

func chMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func EditDistance(a, b string) int {
	dp := make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = 0
	for i := range len(a) + 1 {
		for j := range len(b) + 1 {
			if i > 0 && j > 0 {
				if a[i-1] == b[j-1] {
					dp[i][j] = chMin(dp[i][j], dp[i-1][j-1])
				} else {
					dp[i][j] = chMin(dp[i][j], dp[i-1][j-1]+1)
				}
			}

			if i > 0 {
				dp[i][j] = chMin(dp[i][j], dp[i-1][j]+1)
			}

			if j > 0 {
				dp[i][j] = chMin(dp[i][j], dp[i][j-1]+1)
			}
		}
	}

	return dp[len(a)][len(b)]
}

func SectionDivision(N int, ds [][]int) int {
	dp := make([]int, N+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := range N + 1 {
		for j := range i {
			dp[i] = chMin(dp[i], dp[j]+ds[j][i])
		}
	}
	return dp[N]
}

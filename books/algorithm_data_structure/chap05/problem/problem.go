package problem

import "math"

func chMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Vacation(N int, a, b, c []int) int {
	const CATEGORY = 3
	dp := make([][]int, CATEGORY)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}

	for i := range N + 1 {
		if i > 0 {
			// aをとるケース
			dp[0][i] = chMax(dp[0][i], chMax(dp[1][i-1]+a[i-1], dp[2][i-1]+a[i-1]))
			// bをとるケース
			dp[1][i] = chMax(dp[1][i], chMax(dp[0][i-1]+b[i-1], dp[2][i-1]+b[i-1]))
			// cをとるケース
			dp[2][i] = chMax(dp[2][i], chMax(dp[0][i-1]+c[i-1], dp[1][i-1]+c[i-1]))
		}
	}

	return chMax(dp[0][N], chMax(dp[1][N], dp[2][N]))
}

func IsCreateValueFromAry(W int, ary []int) bool {
	dp := make([][]bool, len(ary)+1)
	for i := range dp {
		dp[i] = make([]bool, W+1)
	}

	dp[0][0] = true
	for i := range ary {
		for j := range W + 1 {
			if dp[i][j] {
				dp[i+1][j] = true
			}

			if j >= ary[i] && dp[i][j-ary[i]] {
				dp[i+1][j] = true
			}
		}
	}

	return dp[len(ary)][W]
}

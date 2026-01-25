package problem

import (
	"math"
)

func Vacation(n int, a, b, c []int) int {
	const CATE = 3
	dp := make([][]int, CATE)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	value := make([][]int, n)
	for i := range value {
		value[i] = make([]int, 3)
		for range value[i] {
			value[i][0] = a[i]
			value[i][1] = b[i]
			value[i][2] = c[i]
		}
	}

	for i := range n {
		for j := range CATE {
			for k := range CATE {
				if j == k {
					continue
				}

				dp[j][i+1] = int(
					math.Max(
						float64(dp[j][i+1]),
						float64(dp[k][i]+value[i][k]),
					),
				)
			}
		}
	}

	var res int
	for i := range CATE {
		res = int(
			math.Max(
				float64(res),
				float64(dp[i][n]),
			),
		)
	}
	return res
}

func SumOfPart(n, w int, a []int) bool {
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true
	for i := range n {
		for j := range w + 1 {
			if dp[i][j] {
				dp[i+1][j] = true
			}

			if j >= a[i] && dp[i][j-a[i]] {
				dp[i+1][j] = true
			}
		}
	}

	return dp[n][w]
}

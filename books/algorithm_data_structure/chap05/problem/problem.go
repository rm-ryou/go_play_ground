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

func ProblemStatement(n int, p []int) int {
	var maxW int
	for i := range p {
		maxW += p[i]
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, maxW+1)
	}

	dp[0][0] = true
	for i := range n {
		for j := range maxW + 1 {
			if dp[i][j] {
				dp[i+1][j] = true
			}

			if j >= p[i] && dp[i][j-p[i]] {
				dp[i+1][j] = true
			}
		}
	}

	var res int
	for i := range dp[n] {
		if dp[n][i] {
			res++
		}
	}
	return res
}

func SumOfPartUnderK(n, w, k int, a []int) bool {
	const inf = int(1 << 60)
	// dp[i][j] aの内先頭からi個選び、jを作る最小の個数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	// aの要素のうち、先頭から0個の要素を用いて0を作れる最小の個数
	dp[0][0] = 0
	for i := range n {
		for j := range w + 1 {
			// 選ばない時
			dp[i+1][j] = int(
				math.Min(
					float64(dp[i+1][j]),
					float64(dp[i][j]),
				),
			)

			// 選ぶ時
			if j >= a[i] {
				dp[i+1][j] = int(
					math.Min(
						float64(dp[i+1][j]),
						float64(dp[i][j-a[i]]+1),
					),
				)
			}
		}
	}

	return dp[n][w] <= k
}

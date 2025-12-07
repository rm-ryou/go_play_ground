package problem

import (
	"math"
)

func chMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func chMin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
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

func Context(ary []int) int {
	var maxValue int
	for i := range ary {
		maxValue += ary[i]
	}
	dp := make([][]bool, len(ary)+1)
	for i := range dp {
		dp[i] = make([]bool, maxValue+1)
	}

	dp[0][0] = true
	for i := range len(ary) {
		for j := range maxValue + 1 {
			if !dp[i][j] {
				continue
			}
			dp[i+1][j] = true
			dp[i+1][j+ary[i]] = true
		}
	}

	res := 0
	for i := range maxValue + 1 {
		if dp[len(ary)][i] {
			res++
		}
	}
	return res
}

func IsCreateValueFromAryUnderK(K, W int, ary []int) bool {
	dp := make([][]int, len(ary)+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
		for j := range dp[i] {
			dp[i][j] = 1e10
		}
	}

	dp[0][0] = 0
	for i := range len(ary) {
		for j := range W + 1 {
			dp[i+1][j] = chMin(dp[i+1][j], dp[i][j])

			if j >= ary[i] {
				dp[i+1][j] = chMin(dp[i+1][j], dp[i][j-ary[i]]+1)
			}
		}
	}

	return dp[len(ary)][W] <= K
}

func IsCreateValueFromAryUnlimited(W int, ary []int) bool {
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

			if j >= ary[i] && dp[i+1][j-ary[i]] {
				dp[i+1][j] = true
			}
		}
	}

	return dp[len(ary)][W]
}

func IsCreateValueFromAryLimited(W int, ary, lim []int) bool {
	dp := make([][]int, len(ary)+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = 0
	for i := range ary {
		for j := range W + 1 {
			dp[i+1][j] = chMin(dp[i+1][j], dp[i][j])

			if j >= ary[i] && dp[i+1][j-ary[i]] <= lim[i] {
				dp[i+1][j] = chMin(dp[i+1][j], dp[i+1][j-ary[i]]+1)
			}
		}
	}

	return dp[len(ary)][W] < math.MaxInt
}

func chMaxStr(s, t string) string {
	if len(s) > len(t) {
		return s
	} else {
		return t
	}
}

// func LCS(s, t string) int {
func LCS(s, t string) string {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	ans := make([][]string, len(s)+1)
	for i := range ans {
		ans[i] = make([]string, len(t)+1)
	}

	for i := range len(s) + 1 {
		for j := range len(t) + 1 {
			if i > 0 && j > 0 {
				if s[i-1] == t[j-1] {
					dp[i][j] = chMax(dp[i][j], dp[i-1][j-1]+1)
					ans[i][j] = ans[i-1][j-1] + string(t[j-1])
				} else {
					dp[i][j] = chMax(dp[i][j], dp[i-1][j-1])
					ans[i][j] = chMaxStr(ans[i][j], ans[i-1][j-1])
				}
			}

			if i > 0 {
				dp[i][j] = chMax(dp[i][j], dp[i-1][j])
				ans[i][j] = chMaxStr(ans[i][j], ans[i-1][j])
			}

			if j > 0 {
				dp[i][j] = chMax(dp[i][j], dp[i][j-1])
				ans[i][j] = chMaxStr(ans[i][j], ans[i][j-1])
			}
		}
	}

	return ans[len(s)][len(t)]
	// return dp[len(s)][len(t)]
}

func SectionDivision(N, M int, ary []int) float64 {
	dp := make([][]float64, N+1)
	for i := range dp {
		dp[i] = make([]float64, M+1)
	}

	aves := make([][]float64, N+1)
	for i := range aves {
		aves[i] = make([]float64, N+1)
	}

	for i := 1; i <= N; i++ {
		for j := 0; j < i; j++ {
			var sum float64
			for k := j; k < i; k++ {
				sum += float64(ary[k])
			}
			aves[j][i] = sum / float64(i-j)
		}
	}

	for i := range N + 1 {
		for j := range i {
			for k := 1; k <= M; k++ {
				dp[i][k] = math.Max(dp[i][k], dp[j][k-1]+aves[j][i])
			}
		}
	}

	var res float64
	for i := range M + 1 {
		res = math.Max(res, dp[N][i])
	}
	return res
}

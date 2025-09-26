package main

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func partOfSumUnderK(n, w, k int, arry []int) bool {
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, w+1)
		for j := range w + 1 {
			dp[i][j] = 1 << 5
		}
	}

	dp[0][0] = 0
	for i := range n {
		for j := range w + 1 {
			dp[i+1][j] = minValue(dp[i+1][j], dp[i][j])

			if j >= arry[i] {
				dp[i+1][j] = minValue(dp[i+1][j], dp[i][j-arry[i]]+1)
			}
		}
	}

	if dp[n][w] <= k {
		return true
	} else {
		return false
	}
}

func partOfSumCounter(n, w int, arry []int) int {
	dp := make([][]bool, n+1)
	for i := range n + 1 {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true
	for i := range n {
		for j := range w + 1 {
			if dp[i][j] {
				dp[i+1][j] = true
			}

			if j >= arry[i] && dp[i][j-arry[i]] {
				dp[i+1][j] = true
			}
		}
	}

	counter := 0
	for i := range w + 1 {
		if dp[n][i] {
			counter++
		}
	}

	return counter - 1
}

func partOfSumBasic(n, w int, arry []int) bool {
	dp := make([][]bool, n+1)
	for i := range n + 1 {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true
	for i := range n {
		for j := range w + 1 {
			if dp[i][j] {
				dp[i+1][j] = true
			}

			if j >= arry[i] && dp[i][j-arry[i]] {
				dp[i+1][j] = true
			}
		}
	}

	return dp[n][w]
}

func vacation(n int, arryA, arryB, arryC []int) int {
	arry := make([][]int, n)
	for i := range n {
		arry[i] = make([]int, 3)
		arry[i][0] = arryA[i]
		arry[i][1] = arryB[i]
		arry[i][2] = arryC[i]
	}
	dpOfVacation := make([][]int, n+1)
	for i := range n + 1 {
		dpOfVacation[i] = make([]int, 3)
	}

	for i := range n {
		for j := range 3 {
			for k := range 3 {
				if j == k {
					continue
				}
				dpOfVacation[i+1][k] = maxValue(dpOfVacation[i+1][k], dpOfVacation[i][j]+arry[i][k])
			}
		}
	}

	return maxValue(dpOfVacation[n][0], maxValue(dpOfVacation[n][1], dpOfVacation[n][2]))
}

func main() {}

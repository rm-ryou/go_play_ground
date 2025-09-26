package main

import (
	"fmt"
	"math"
)

func getMin(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	S, T := "logistic", "algorithm"

	memo := make([][]int, len(S)+1)
	for i := range len(S) + 1 {
		memo[i] = make([]int, len(T)+1)
		for j := range len(memo[i]) {
			memo[i][j] = math.MaxInt
		}
	}

	for i := 0; i <= len(S); i++ {
		for j := 0; j <= len(T); j++ {
			if i > 0 && j > 0 {
				if S[i-1] == T[j-1] {
					memo[i][j] = getMin(memo[i][j], memo[i-1][j-1])
				} else {
					memo[i][j] = getMin(memo[i][j], memo[i-1][j-1]+1)
				}
			}

			if i > 0 {
				memo[i][j] = getMin(memo[i][j], memo[i-1][j]+1)
			}

			if j > 0 {
				memo[i][j] = getMin(memo[i][j], memo[i][j-1]+1)
			}
		}
	}

	fmt.Println(memo[len(S)][len(T)])
}

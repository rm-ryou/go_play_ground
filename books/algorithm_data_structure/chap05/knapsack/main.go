package main

import (
	"fmt"
	"math"
)

func main() {
	n, w := 6, 9
	vals := []int{3, 2, 6, 1, 3, 85}
	weights := []int{2, 1, 3, 2, 1, 5}

	memo := make([][]int, n+1)
	for i := range n + 1 {
		memo[i] = make([]int, w+1)
	}

	for i := range n {
		for j := range w + 1 {
			if j-weights[i] >= 0 {
				maxV := math.Max(
					float64(memo[i+1][j]),
					float64(memo[i][j-weights[i]]+vals[i]),
				)
				memo[i+1][j] = int(maxV)
			}

			maxV := math.Max(float64(memo[i+1][j]), float64(memo[i][j]))
			memo[i+1][j] = int(maxV)
		}
	}

	fmt.Println(memo[n][w])
}

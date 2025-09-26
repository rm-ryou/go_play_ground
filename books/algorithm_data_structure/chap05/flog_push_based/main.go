package main

import (
	"fmt"
	"math"
)

func main() {
	n := 7
	arry := []int{2, 9, 4, 5, 1, 6, 10}
	memo := make([]int, n+1)
	for i := range len(memo) {
		memo[i] = math.MaxInt
	}

	memo[0] = 0

	for i := range n {
		if i+1 < n {
			minV := math.Min(
				float64(memo[i+1]),
				float64(memo[i]+int(math.Abs(float64(arry[i]-arry[i+1])))),
			)
			memo[i+1] = int(minV)
		}
		if i+2 < n {
			minV := math.Min(
				float64(memo[i+2]),
				float64(memo[i]+int(math.Abs(float64(arry[i]-arry[i+2])))),
			)
			memo[i+1] = int(minV)
		}
	}

	fmt.Println(memo)
	fmt.Println(memo[n-1])
}

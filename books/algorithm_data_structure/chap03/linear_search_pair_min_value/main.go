package main

import (
	"fmt"
	"math"
)

func searchPairMinValue(a, b []int, k int) int {
	res := math.MaxInt

	for _, va := range a {
		for _, vb := range b {
			if va+vb < k {
				continue
			}

			if va+vb < res {
				res = va + vb
			}
		}
	}
	return res
}

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	arryA := make([]int, N)
	for i := range N {
		var n int
		fmt.Scanf("%d", &n)
		arryA[i] = n
	}

	arryB := make([]int, N)
	for i := range N {
		var n int
		fmt.Scanf("%d", &n)
		arryB[i] = n
	}

	res := searchPairMinValue(arryA, arryB, K)
	fmt.Println(res)
}

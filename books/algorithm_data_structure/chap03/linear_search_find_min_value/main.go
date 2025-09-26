package main

import (
	"fmt"
	"math"
)

func findMinValue(arry []int) int {
	res := math.MaxInt

	for _, v := range arry {
		if v < res {
			res = v
		}
	}

	return res
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	arry := make([]int, N)
	for i := range N {
		var n int
		fmt.Scanf("%d", &n)
		arry[i] = n
	}

	fmt.Println(findMinValue(arry))
}

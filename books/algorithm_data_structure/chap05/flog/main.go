package main

import (
	"fmt"
	"math"
)

var memo []int

func flog(curPosition int, arry []int) {
	if curPosition >= len(arry) {
		return
	}

	if curPosition == 1 {
		memo[curPosition] = int(math.Abs(float64(arry[curPosition] - arry[curPosition-1])))
		flog(curPosition+1, arry)
	} else {
		memo[curPosition] = int(
			math.Min(
				float64(memo[curPosition-1]+int(math.Abs(float64(arry[curPosition]-arry[curPosition-1])))),
				float64(memo[curPosition-1]+int(math.Abs(float64(arry[curPosition]-arry[curPosition-2])))),
			),
		)
		flog(curPosition+1, arry)
	}
}

func main() {
	n := 7
	arry := []int{2, 9, 4, 5, 1, 6, 10}
	memo = make([]int, n+1)
	for i := range len(memo) {
		memo[i] = math.MaxInt
	}

	memo[0] = 0
	flog(1, arry)

	fmt.Println(memo)
	fmt.Println(memo[n-1])
}

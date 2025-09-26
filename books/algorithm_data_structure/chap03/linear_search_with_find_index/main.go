package main

import "fmt"

func linearSearchWithIndex(arry []int, val int) int {
	res := -1
	for i, n := range arry {
		if n == val {
			res = i
		}
	}
	return res
}

func main() {
	var N, v int
	fmt.Scanf("%d %d", &N, &v)

	arry := make([]int, N)
	for i := range N {
		var n int
		fmt.Scanf("%d", &n)
		arry[i] = n
	}

	fmt.Println(linearSearchWithIndex(arry, v))
}

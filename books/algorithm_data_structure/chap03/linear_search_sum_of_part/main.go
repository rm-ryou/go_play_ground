package main

import "fmt"

func canCreateW(arry []int, w int) bool {
	res := false

	for bit := range 1 << len(arry) {
		sum := 0
		for i, v := range arry {
			if bit&(1<<i) > 0 {
				sum += v
			}
		}

		if sum == w {
			res = true
		}
	}
	return res
}

func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)

	arry := make([]int, N)
	for i := range N {
		var n int
		fmt.Scanf("%d", &n)
		arry[i] = n
	}

	if canCreateW(arry, W) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import "fmt"

func linearSearch(arry []int, v int) bool {
	res := false

	for _, num := range arry {
		if num == v {
			res = true
		}
	}
	return res
}

func main() {
	var N, v int
	fmt.Scanf("%d %v", &N, &v)

	arry := make([]int, N)
	for i := range N {
		var num int
		fmt.Scanf("%d", &num)
		arry[i] = v
	}

	if linearSearch(arry, v) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

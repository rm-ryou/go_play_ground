package main

import "fmt"

var memo []int

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fibo(n-1) + fibo(n-2)
	return memo[n]
}

func main() {
	n := 49
	memo = make([]int, n+1)

	fmt.Println(fibo(n))
}

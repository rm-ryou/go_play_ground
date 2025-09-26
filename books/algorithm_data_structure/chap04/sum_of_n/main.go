package main

import "fmt"

func fn(n int) int {
	if n == 0 {
		return 0
	}
	return n + fn(n-1)
}

func main() {
	N := 10

	fmt.Println(fn(N))
}

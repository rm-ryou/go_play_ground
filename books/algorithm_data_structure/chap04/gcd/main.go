package main

import "fmt"

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func main() {
	fmt.Println(gcd(51, 15))
	fmt.Println(gcd(15, 51))
}

package main

import "fmt"

func multiply(value, multiplier int) int {
	return value * multiplier
}

func add(value, additive int) int {
	return value + additive
}

func main() {
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}

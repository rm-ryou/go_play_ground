package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func What() func(int) func(int) func(int) int {
	sum := 0
	return func(x int) func(int) func(int) int {
		sum += x
		return func(y int) func(int) int {
			sum += y
			return func(z int) int {
				sum += z
				return sum
			}
		}
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	what := What()
	for i := 0; i < 5; i++ {
		fmt.Println(what(i)(i + 1)(i + 2))
	}
	tmp := what(1)(2)
	fmt.Println(tmp(3))
}

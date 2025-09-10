package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func compute_v2(x, y float64, fn func(float64, float64) float64) float64 {
	return (fn(x, y))
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute_v2(2, 2, hypot))
	fmt.Println(compute(math.Pow))
}

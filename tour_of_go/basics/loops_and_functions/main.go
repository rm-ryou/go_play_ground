package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= float64((z*z - x) / (2 * z))
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("math.Sqrt(): ", math.Sqrt(2))
	fmt.Println("#############")
	fmt.Println(Sqrt(3))
	fmt.Println("math.Sqrt(): ", math.Sqrt(3))
	fmt.Println("#############")
	fmt.Println(Sqrt(4))
	fmt.Println("math.Sqrt(): ", math.Sqrt(4))
	fmt.Println("#############")
	fmt.Println(Sqrt(10))
	fmt.Println("math.Sqrt(): ", math.Sqrt(10))
}

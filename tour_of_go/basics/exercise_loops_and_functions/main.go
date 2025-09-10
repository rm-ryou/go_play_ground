package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	res := 1.0
	tmp := res
	for i := 0; i < 10; i++ {
		tmp = res
		res -= (res*res - x) / (2 * res)
		if tmp == res {
			fmt.Printf("Not change!! index: %d, val: %v\n", i, res)
			break
		}
	}
	if tmp != res {
		fmt.Printf("diff: %v %v\n", res, tmp)
	}
	return res
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(4), Sqrt(4))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(5), Sqrt(5))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(6), Sqrt(6))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(42), Sqrt(42))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(142), Sqrt(142))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(242), Sqrt(242))
	fmt.Printf("math pkg: %v, custom: %v\n", math.Sqrt(999), Sqrt(999))
}

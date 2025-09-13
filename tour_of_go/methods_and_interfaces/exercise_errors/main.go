package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	res := 1.0
	for range 10 {
		res -= (res*res - x) / (2 * res)
	}
	return res, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	num := float64(e)
	return fmt.Sprintf("Can't Sqrt negative number: %f", num)
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	for range 10 {
		z -= (z*z - x) / (2 * z)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

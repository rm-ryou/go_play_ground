package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	num := float64(e)
	return fmt.Sprintf("can't Sqrt negative number: %v", num)
}

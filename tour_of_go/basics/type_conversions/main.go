package main

import (
	"fmt"
	"math"
)

func main() {
	i := 42
	j := float64(i)
	var k uint = uint(j)
	fmt.Printf("Value: %v Type: %T\n", i, i)
	fmt.Printf("Value: %v Type: %T\n", j, j)
	fmt.Printf("Value: %v Type: %T\n", k, k)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

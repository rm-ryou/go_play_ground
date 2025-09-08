package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)

	a := 65
	fmt.Println(a)
	fmt.Printf("convert2string: %q\n", string(a))
	fmt.Printf("convert2rune: %c\n", rune(a))
}

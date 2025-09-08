package main

import (
	"fmt"
	"math"

	"example.com/pkg"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(pkg.Hoge())
}

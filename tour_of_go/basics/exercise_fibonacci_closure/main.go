package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		res := x
		x, y = y, x+y
		return res
	}
	// x, y := 0, 0
	// return func() int {
	// 	x, y = y, x+y
	// 	if y == 0 {
	// 		y += 1
	// 	}
	// 	return x
	// }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	var i int
	j := i
	fmt.Println(i == j)

	f := 3.14
	fmt.Printf("v is of type %T\n", f)
	g := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", g)

	// u := 1<<64 - 1
	// fmt.Printf("v is of type %T\n", u)
}

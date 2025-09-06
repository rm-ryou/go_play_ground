package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	i := 42
	f := 3.142
	g := 0.867 + 0.5i
	fmt.Printf("Type: %T, Value: %v\n", i, i)
	fmt.Printf("Type: %T, Value: %v\n", f, f)
	fmt.Printf("Type: %T, Value: %v\n", g, g)

	h := "Hello, World!"
	c := 'A'
	fmt.Printf("Type: %T, Value: %v\n", h, h)
	fmt.Printf("Type: %T, Value: %v\n", c, c)
}

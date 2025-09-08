package main

import "fmt"

func main() {
	s := "string"
	i, j := 1, 2
	c, python, java := true, false, "no!"

	vals := []any{s, i, j, c, python, java}
	for _, v := range vals {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

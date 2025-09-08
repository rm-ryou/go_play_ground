package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	vals := []any{i, j, c, python, java}
	for _, v := range vals {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

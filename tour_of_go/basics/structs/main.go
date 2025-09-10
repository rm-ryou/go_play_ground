package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type pVertex struct {
	X, Y *int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{})

	fmt.Println(pVertex{})
}

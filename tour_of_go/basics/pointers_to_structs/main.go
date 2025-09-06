package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z *int
}

func main() {
	num := 42
	v := Vertex{X: 1, Y: 2, Z: &num}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v.Z)
	fmt.Println(*(v.Z))
}

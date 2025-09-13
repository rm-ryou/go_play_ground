package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return v.X*v.X + v.Y*v.Y
}

func main() {
	v := Vertex{
		X: 3,
		Y: 4,
	}
	fmt.Println(Abs(v))
}

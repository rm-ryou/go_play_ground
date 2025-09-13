package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{
		X: 3,
		Y: 4,
	}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{
		X: 4,
		Y: 3,
	}
	fmt.Println((*p).Abs())
	fmt.Println(AbsFunc(*p))
}

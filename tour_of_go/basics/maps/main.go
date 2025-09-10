package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	fmt.Println(m)
	if m == nil {
		println("m is nil")
	}

	m = make(map[string]Vertex)
	fmt.Println(m)
	if m == nil {
		println("m is nil")
	}
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

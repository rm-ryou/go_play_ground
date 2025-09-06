package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	empty := make(map[int]int, 0)
	fmt.Println(empty)
	fmt.Println(empty[1])
	empty[1] = 1
	fmt.Println(empty[1])
}

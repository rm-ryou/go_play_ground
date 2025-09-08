package main

import "fmt"

const (
	Pi = 3.14
	pi = 3.1415
)

func main() {
	const world = "World!"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Happy", pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)
}

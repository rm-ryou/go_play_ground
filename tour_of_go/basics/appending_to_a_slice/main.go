package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	cp := make([]int, len(s))
	n := copy(cp, s)
	printSlice(cp)
	fmt.Println(n)

	newSlice := append(cp, s...)
	printSlice(newSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

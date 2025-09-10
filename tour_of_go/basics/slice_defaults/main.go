package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len: %d cap: %d value: %v\n", len(s), cap(s), s)
	// 6, 6, ...

	s = s[1:4]
	fmt.Printf("len: %d cap: %d value: %v\n", len(s), cap(s), s)
	// 3, 5, 3 5 7

	s = s[:2]
	fmt.Printf("len: %d cap: %d value: %v\n", len(s), cap(s), s)
	// 2, 5, 3 5

	s = s[1:]
	fmt.Printf("len: %d cap: %d value: %v\n", len(s), cap(s), s)
	// 1, 4, 5
}

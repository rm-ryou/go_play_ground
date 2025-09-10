package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s, len(s), cap(s))
	s = s[1:4]
	fmt.Println(s, len(s), cap(s))
	s = s[0:4]
	fmt.Println(s, len(s), cap(s))
	s = s[2:4]
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s[:])
}

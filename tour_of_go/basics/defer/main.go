package main

import "fmt"

func main() {
	defer fmt.Println("defer")
	defer fmt.Println("hello")
	defer fmt.Println("world")

	fmt.Println("Hello")
	defer fmt.Println("last line")
}

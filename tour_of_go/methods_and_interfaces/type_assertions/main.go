package main

import "fmt"

func main() {
	var i any = "hello"
	fmt.Println(i)

	s := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

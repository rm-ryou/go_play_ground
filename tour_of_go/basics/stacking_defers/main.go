package main

import "fmt"

func main() {
	fmt.Println("counting")

	var i int
	for i = 0; i < 10; i++ {
		defer fmt.Println(i)
		defer func() {
			i += 1
			fmt.Println(i)
		}()
	}
	fmt.Printf("i = %d\n", i)

	fmt.Println("done")
}

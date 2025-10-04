package main

import "fmt"

func main() {
	intCh := make(chan int, 0)
	go func() {
		intCh <- 1
	}()
	fmt.Println(<-intCh)
}

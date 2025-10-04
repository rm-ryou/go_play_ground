package main

import "fmt"

func main() {
	intCh := make(chan int)
	close(intCh)
	integer, ok := <-intCh
	fmt.Printf("(%v): %v\n", ok, integer)

	hogeCh := make(chan int)
	go func() {
		defer close(hogeCh)
		for i := range 5 {
			hogeCh <- i
		}
	}()

	for i := range hogeCh {
		fmt.Printf("%v ", i)
	}
}

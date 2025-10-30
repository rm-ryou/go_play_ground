package main

import "fmt"

func chanOwner() <-chan int {
	resCh := make(chan int, 5)
	go func() {
		defer close(resCh)
		for i := range 5 {
			resCh <- i
		}
	}()
	return resCh
}

func main() {
	for i := range chanOwner() {
		fmt.Println(i)
	}
}

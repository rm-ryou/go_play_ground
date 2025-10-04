package main

import "fmt"

func chanOwner() <-chan int {
	res := make(chan int, 5)
	go func() {
		defer close(res)
		for i := range 5 {
			res <- i
		}
	}()
	return res
}

func consumer(ch <-chan int) {
	for res := range ch {
		fmt.Printf("Received: %d\n", res)
	}
	fmt.Println("Done receiving!")
}

func main() {
	res := chanOwner()
	consumer(res)
}

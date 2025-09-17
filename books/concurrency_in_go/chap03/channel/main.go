package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		resultCh := make(chan int, 5)
		go func() {
			defer close(resultCh)
			for i := range 5 {
				resultCh <- i
			}
		}()
		return resultCh
	}

	resultCh := chanOwner()
	for res := range resultCh {
		fmt.Printf("Received: %d\n", res)
	}
	fmt.Println("Done receiving!")
}

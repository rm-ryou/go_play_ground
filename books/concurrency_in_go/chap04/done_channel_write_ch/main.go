package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randCh(done <-chan any) <-chan int {
	ch := make(chan int)

	go func() {
		defer fmt.Println("done randCh")
		defer close(ch)
		for {
			select {
			case <-done:
				return
				// default:
			case ch <- rand.Int():
			}
			// ch <- rand.Int()
		}
	}()

	return ch
}

func main() {
	done := make(chan any)
	ch := randCh(done)

	for range 3 {
		fmt.Printf("%d\n", <-ch)
	}
	close(done)
	time.Sleep(1 * time.Second)
}

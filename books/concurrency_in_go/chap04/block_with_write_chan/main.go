package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newRandCh() <-chan int {
	randCh := make(chan int)
	go func() {
		defer fmt.Println("newRandCh closure exited.")
		defer close(randCh)
		for {
			randCh <- rand.Int()
		}
	}()
	return randCh
}

func newRandChWithCancel(done <-chan int) <-chan int {
	randCh := make(chan int)
	go func() {
		defer fmt.Println("newRandChWithCancel closure exited.")
		defer close(randCh)
		for {
			select {
			case randCh <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return randCh
}

func main() {
	// randCh := newRandCh()
	// fmt.Println("3 random ints:")
	// for i := range 3 {
	// 	fmt.Printf("%d: %d\n", i, <-randCh)
	// }

	done := make(chan int)
	randCh := newRandChWithCancel(done)
	fmt.Println("3 random ints:")
	for i := range 3 {
		fmt.Printf("%d: %d\n", i, <-randCh)
	}
	close(done)
	time.Sleep(1 * time.Second)
}

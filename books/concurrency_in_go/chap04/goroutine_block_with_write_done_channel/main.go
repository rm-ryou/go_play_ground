package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandCh := func(done <-chan any) <-chan int {
		randCh := make(chan int)
		go func() {
			defer fmt.Println("newRandCh closure exited.")
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

	done := make(chan any)
	randCh := newRandCh(done)
	fmt.Println("3 random ints:")
	for i := range 3 {
		fmt.Printf("%d: %d\n", i, <-randCh)
	}
	close(done)
	time.Sleep(1 * time.Second)
}

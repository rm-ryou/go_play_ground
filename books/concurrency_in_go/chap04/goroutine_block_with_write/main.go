package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandCh := func() <-chan int {
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

	randCh := newRandCh()
	fmt.Println("3 random ints:")
	for i := range 3 {
		fmt.Printf("%d: %d\n", i, <-randCh)
	}

	time.Sleep(1 * time.Second)
}

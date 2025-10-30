package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan any)
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()

	counter := 0
loop:
	for {
		select {
		case <-ch:
			break loop
		default:
		}

		counter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achived %v cycles of work before signaled to stop.\n", counter)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan any)
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	var workCounter int
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Archived %v cycles of work before signalled to stop\n", workCounter)
}

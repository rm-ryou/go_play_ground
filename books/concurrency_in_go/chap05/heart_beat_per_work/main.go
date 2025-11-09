package main

import (
	"fmt"
	"math/rand"
)

func doWork(done <-chan any) (<-chan any, <-chan int) {
	heartbeatCh := make(chan any)
	workCh := make(chan int)

	go func() {
		defer close(heartbeatCh)
		defer close(workCh)

		for range 10 {
			select {
			case heartbeatCh <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case workCh <- rand.Intn(10):
			}
		}
	}()

	return heartbeatCh, workCh
}

func main() {
	done := make(chan any)
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("results %v\n", r)
		}
	}
}

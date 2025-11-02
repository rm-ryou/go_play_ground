package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan any, pulseInterval time.Duration) (<-chan any, <-chan time.Time) {
	heartBeat := make(chan any)
	results := make(chan time.Time)

	go func() {
		// defer close(heartBeat)
		// defer close(results)

		pulse := time.Tick(pulseInterval)
		genWork := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartBeat <- struct{}{}:
			default:
			}
		}

		sendResults := func(r time.Time) {
			for {
				select {
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for range 2 {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-genWork:
				sendResults(r)
			}
		}
	}()

	return heartBeat, results
}

func main() {
	done := make(chan any)
	time.AfterFunc(10*time.Second, func() { close(done) })

	const timeout = 2 * time.Second
	heartBeat, results := doWork(done, timeout/2)

	for {
		select {
		case _, ok := <-heartBeat:
			if !ok {
				return
			}
			fmt.Println("Pulse!!")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("results %v\n", r)
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}

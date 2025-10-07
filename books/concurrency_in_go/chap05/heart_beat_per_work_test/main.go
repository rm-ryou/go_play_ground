package main

import "time"

func doWork(done <-chan any, nums ...int) (<-chan any, <-chan int) {
	heartBeat := make(chan any, 1)
	intCh := make(chan int)

	go func() {
		defer close(heartBeat)
		defer close(intCh)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartBeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case intCh <- n:
			}
		}
	}()

	return heartBeat, intCh
}

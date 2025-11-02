package main

import "sync"

func fanIn(done <-chan any, channels ...<-chan any) <-chan any {
	ch := make(chan any)
	var wg sync.WaitGroup

	multiplex := func(c <-chan any) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

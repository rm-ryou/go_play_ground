package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(done <-chan any, id int, wg *sync.WaitGroup, res chan<- int) {
	started := time.Now()
	defer wg.Done()

	simulatedLoadTime := time.Duration(1*rand.Intn(5)) * time.Second
	select {
	case <-done:
	case <-time.After(simulatedLoadTime):
	}

	select {
	case <-done:
	case res <- id:
	}

	took := time.Since(started)
	if took < simulatedLoadTime {
		took = simulatedLoadTime
	}
	fmt.Printf("%v took %v\n", id, took)
}

func main() {
	done := make(chan any)
	res := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := range 10 {
		go doWork(done, i, &wg, res)
	}

	firstReturned := <-res
	close(done)
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)
}

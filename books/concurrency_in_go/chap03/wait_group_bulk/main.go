package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v!\n", id)
}

func main() {
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := range numGreeters {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

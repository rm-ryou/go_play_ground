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
	for i := range numGreeters {
		wg.Add(1)
		go hello(&wg, i+1)
	}
	wg.Wait()
}

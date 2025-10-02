package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}

func decrement() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Printf("Decrementing: %d\n", count)
}

func main() {
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("Completed")
}

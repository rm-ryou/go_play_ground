package main

import (
	"fmt"
	"sync"
)

var (
	count int
	lock  sync.Mutex
)

func Inc() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}

func Dec() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Printf("Decrementing: %d\n", count)
}

func main() {
	var arithmetic sync.WaitGroup
	for range 5 {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			Inc()
		}()
	}

	for range 5 {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			Dec()
		}()
	}
	arithmetic.Wait()

	fmt.Println("Arithmetic complete.")
}

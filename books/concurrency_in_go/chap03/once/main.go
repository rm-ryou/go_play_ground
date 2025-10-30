package main

import (
	"fmt"
	"sync"
)

var count = 0

func increment() {
	count++
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	n := 100
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()

	fmt.Println("count is:", count)
}

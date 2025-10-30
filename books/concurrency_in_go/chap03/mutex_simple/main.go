package main

import (
	"fmt"
	"sync"
)

type Count struct {
	mu  sync.Mutex
	val int
}

func (c *Count) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
	fmt.Printf("Incrementing: %d\n", c.val)
}

func (c *Count) decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val--
	fmt.Printf("Decrementing: %d\n", c.val)
}

func main() {
	count := &Count{val: 0}

	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.increment()
		}()
	}

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.decrement()
		}()
	}

	wg.Wait()
	fmt.Println("complete count is:", count.val)
}

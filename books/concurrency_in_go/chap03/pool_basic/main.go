package main

import (
	"fmt"
	"sync"
)

func main() {
	var numMounted int
	pool := sync.Pool{
		New: func() any {
			numMounted++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())

	var wg sync.WaitGroup
	const numWorkers = 1024 * 1024

	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			mem := pool.Get().(*[]byte)
			defer pool.Put(mem)

		}()
	}
	wg.Wait()

	fmt.Println(numMounted)
}

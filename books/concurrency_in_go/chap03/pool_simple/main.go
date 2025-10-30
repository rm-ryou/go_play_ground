package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	pool.Get()
	instance := pool.Get()
	pool.Put(instance)
	pool.Get()
}

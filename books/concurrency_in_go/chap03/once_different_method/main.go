package main

import (
	"fmt"
	"sync"
)

var count = 0

func increment() {
	count++
}

func decrement() {
	count--
}

func main() {
	var once sync.Once

	once.Do(increment)
	once.Do(decrement)

	fmt.Println("count is:", count)
}

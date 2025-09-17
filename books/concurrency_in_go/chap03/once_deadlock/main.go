package main

import (
	"fmt"
	"sync"
)

func main() {
	var onceA, onceB sync.Once

	var initB func()
	initA := func() {
		fmt.Println("InitA!!")
		onceB.Do(initB)
	}
	initB = func() {
		fmt.Println("InitB!!")
		onceA.Do(initA)
	}
	fmt.Println("process!")
	onceA.Do(initA)
}

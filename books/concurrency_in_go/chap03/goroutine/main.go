package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()

	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "Welcome!"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

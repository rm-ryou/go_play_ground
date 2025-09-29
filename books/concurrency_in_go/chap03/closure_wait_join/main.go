package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	message := "hello!"

	wg.Add(1)
	go func() {
		defer wg.Done()
		message = "welcome!"
	}()

	fmt.Println(message)
	wg.Wait()
	fmt.Println(message)

	printSalutation()
}

func printSalutation() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

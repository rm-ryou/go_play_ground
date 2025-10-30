package main

import (
	"fmt"
)

func main() {
	stringCh := make(chan string, 0)

	go func() {
		stringCh <- "Hello!!"
	}()

	fmt.Println(<-stringCh)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	select {
	case <-ch:
	case <-time.After(1 * time.Second):
		fmt.Println("time outed")
	}
}

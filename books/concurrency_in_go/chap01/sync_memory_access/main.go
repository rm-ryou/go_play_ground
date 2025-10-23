package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var data int

	go func() {
		memoryAccess.Lock()
		defer memoryAccess.Unlock()
		data++
	}()

	memoryAccess.Lock()
	defer memoryAccess.Unlock()
	if data == 0 {
		fmt.Println("the value is 0.\n")
	} else {
		fmt.Printf("the value is %v\n", data)
	}
}

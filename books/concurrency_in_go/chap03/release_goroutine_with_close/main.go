package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan any)
	var wg sync.WaitGroup
	for i := range 4 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	time.Sleep(1 * time.Second)
	close(begin)
	wg.Wait()
}

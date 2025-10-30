package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan any)
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-begin
			fmt.Printf("%d has begun\n", i)
		}()
	}

	fmt.Println("unblocking goroutines...")
	time.Sleep(2 * time.Second)
	close(begin)
	wg.Wait()
}

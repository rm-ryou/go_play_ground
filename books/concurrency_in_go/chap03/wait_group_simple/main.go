package main

import (
	"fmt"
	"sync"
	"time"
)

func bulkWait(wg *sync.WaitGroup, n int) {
	start := time.Now()
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
		}()
		// go func(i int) {
		// 	defer wg.Done()
		// 	fmt.Println(i)
		// }(i)
	}
	wg.Wait()
	fmt.Println("bulkWait time is:", time.Since(start))
}

func eachWait(wg *sync.WaitGroup, n int) {
	start := time.Now()
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		// go func(i int) {
		// 	defer wg.Done()
		// 	fmt.Println(i)
		// }(i)
	}
	wg.Wait()
	fmt.Println("eachWait time is:", time.Since(start))
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		// time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		// time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete")

	bulkWait(&wg, 100)
	eachWait(&wg, 100)
}

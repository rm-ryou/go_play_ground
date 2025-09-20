package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan any {
		completed := make(chan any)
		go func() {
			fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// ch := doWork(nil)
	// <-ch
	fmt.Println("Done.")
}

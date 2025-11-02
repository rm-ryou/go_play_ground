package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan any, strings <-chan string) <-chan any {
	res := make(chan any)
	go func() {
		defer fmt.Println("doWork is done.")
		defer close(res)
		for {
			select {
			case <-done:
				return
			case s := <-strings:
				fmt.Println(s)
			}
		}
	}()
	return res
}

func main() {
	done := make(chan any)
	res := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	<-res
	fmt.Println("done!")
}

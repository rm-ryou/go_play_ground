package main

import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)

	var count1, count2 int
	for range 1000 {
		select {
		case <-ch1:
			count1++
		case <-ch2:
			count2++
		}
	}

	fmt.Printf("count1: %d\ncount2: %d\n", count1, count2)
}

package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for range n {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// NOTE: チャネルが閉じられるまで繰り返し受信し続ける
	// そのため、送り手で明示的にチャネルをcloseする必要がある
	for i := range c {
		fmt.Println(i)
	}
}

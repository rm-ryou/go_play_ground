package main

import "fmt"

func repeat(done <-chan any, values ...any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)

		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case ch <- v:
				}
			}
		}
	}()

	return ch
}

func take(done <-chan any, valueCh <-chan any, num int) <-chan any {
	takeCh := make(chan any)
	go func() {
		defer close(takeCh)
		for range num {
			select {
			case <-done:
				return
			case takeCh <- <-valueCh:
			}
		}
	}()

	return takeCh
}

func main() {
	done := make(chan any)
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}

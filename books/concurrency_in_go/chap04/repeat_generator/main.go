package main

import "fmt"

func repeat(done <-chan any, values ...any) <-chan any {
	valueCh := make(chan any)
	go func() {
		defer close(valueCh)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueCh <- v:
				}
			}
		}
	}()
	return valueCh
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

	for num := range take(done, repeat(done, 1, "hello"), 10) {
		fmt.Printf("%v ", num)
	}
}

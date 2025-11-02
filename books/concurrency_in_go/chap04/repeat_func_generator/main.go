package main

import (
	"fmt"
	"math/rand"
)

func repeatFn(done <-chan any, fn func() any) <-chan any {
	repeatCh := make(chan any)
	go func() {
		defer close(repeatCh)
		for {
			select {
			case <-done:
				return
			case repeatCh <- fn():
			}
		}
	}()

	return repeatCh
}

func take(done <-chan any, valueCh <-chan any, num int) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)

		for range num {
			select {
			case <-done:
				return
			case ch <- <-valueCh:
			}
		}
	}()

	return ch
}

func main() {
	done := make(chan any)
	defer close(done)

	rand := func() any { return rand.Int() }
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

package main

import (
	"fmt"
	"math/rand"
)

func repeat(done <-chan any, fn func() any) <-chan any {
	valueCh := make(chan any)
	go func() {
		defer close(valueCh)
		for {
			select {
			case <-done:
				return
			case valueCh <- fn():
			}
		}
	}()
	return valueCh
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
	for num := range take(done, repeat(done, rand), 10) {
		fmt.Println(num)
	}
}

package main

import (
	"fmt"
	"math/rand"
)

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

	for num := range take(done, repeat(done, 1, 2, "hoge"), 10) {
		fmt.Printf("%v\n", num)
	}

	rand := func() any { return rand.Int() }
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

func repeatFn(done <-chan any, fn func() any) <-chan any {
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

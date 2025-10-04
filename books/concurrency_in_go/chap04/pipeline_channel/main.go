package main

import "fmt"

func generator(done <-chan any, integers ...int) <-chan int {
	intCh := make(chan int, len(integers))
	go func() {
		defer close(intCh)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intCh <- i:
			}
		}
	}()
	return intCh
}

func multiply(done <-chan any, intCh <-chan int, multiplier int) <-chan int {
	multipliedCh := make(chan int)
	go func() {
		defer close(multipliedCh)
		for i := range intCh {
			select {
			case <-done:
				return
			case multipliedCh <- i * multiplier:
			}
		}
	}()
	return multipliedCh
}

func add(done <-chan any, intCh <-chan int, additive int) <-chan int {
	addedCh := make(chan int)
	go func() {
		defer close(addedCh)
		for i := range intCh {
			select {
			case <-done:
				return
			case addedCh <- i + additive:
			}
		}
	}()
	return addedCh
}

func main() {
	done := make(chan any)
	defer close(done)

	intCh := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intCh, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}

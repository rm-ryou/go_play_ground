package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFn(done <-chan any, fn func() any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- fn():
			}
		}
	}()
	return ch
}

func toInt(done <-chan any, valueCh <-chan any) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range valueCh {
			select {
			case <-done:
				return
			case ch <- v.(int):
			}
		}
	}()
	return ch
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

func primeFinder(done <-chan any, intCh <-chan int) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for i := range intCh {
			i -= 1
			prime := true
			for divisor := i - 1; divisor > 1; divisor-- {
				if i%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case ch <- i:
				}
			}
		}
	}()
	return ch
}

func fanIn(done <-chan any, channels ...<-chan any) <-chan any {
	var wg sync.WaitGroup
	multiplexedCh := make(chan any)

	multiplex := func(c <-chan any) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedCh <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedCh)
	}()
	return multiplexedCh
}

func main() {
	rand := func() any { return rand.Intn(50000000) }

	done := make(chan any)
	defer close(done)

	start := time.Now()

	// randIntCh := toInt(done, repeatFn(done, rand))
	// fmt.Println("Primes:")
	// for prime := range take(done, primeFinder(done, randIntCh), 10) {
	// 	fmt.Printf("\t%d\n", prime)
	// }

	randIntCh := toInt(done, repeatFn(done, rand))
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan any, numFinders)
	fmt.Println("Primes:")
	for i := range numFinders {
		finders[i] = primeFinder(done, randIntCh)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

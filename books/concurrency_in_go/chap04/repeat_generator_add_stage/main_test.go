package main

import "testing"

func BenchmarkGeneric(b *testing.B) {
	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	repeat := func(done <-chan any, values ...string) <-chan string {
		valueCh := make(chan string)
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

	take := func(done <-chan any, valueCh <-chan string, num int) <-chan string {
		takeCh := make(chan string)
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

	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for range take(done, repeat(done, "a"), b.N) {
	}
}

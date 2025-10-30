package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup

	begin := make(chan any)
	c := make(chan any)

	var token any
	sender := func() {
		defer wg.Done()
		<-begin
		for range b.N {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for range b.N {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

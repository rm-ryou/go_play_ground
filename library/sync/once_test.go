// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	stdSync "sync"
	"sync/atomic"
	"testing"
)

func TestOnce(t *testing.T) {
	var once Once
	var x int
	for i := 0; i < 10; i++ {
		once.Do(func() {
			x++
		})
	}
	if x != 1 {
		t.Fatalf("once.Do called function %d times, want 1", x)
	}
}

func TestOnceConcurrent(t *testing.T) {
	var once Once
	var calls int32
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				atomic.AddInt32(&calls, 1)
			})
			c <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-c
	}

	if calls != 1 {
		t.Fatalf("once.Do called function %d times, want 1", calls)
	}
}

func TestOncePanic(t *testing.T) {
	var once Once
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("once.Do did not panic")
			}
		}()
		once.Do(func() {
			panic("failed")
		})
	}()

	called := false
	once.Do(func() {
		called = true
	})
	if called {
		t.Fatal("once.Do called function after panic")
	}
}

func TestOnceBlocksUntilDone(t *testing.T) {
	var once Once
	started := make(chan bool)
	release := make(chan bool)
	done := make(chan bool)

	go func() {
		once.Do(func() {
			started <- true
			<-release
		})
	}()
	<-started

	go func() {
		once.Do(func() {
			t.Error("once.Do called function twice")
		})
		done <- true
	}()

	select {
	case <-done:
		t.Fatal("once.Do returned before first function completed")
	default:
	}

	release <- true
	<-done
}

func BenchmarkStdOnce(b *testing.B) {
	var once stdSync.Once
	f := func() {}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(f)
		}
	})
}

func BenchmarkOnce(b *testing.B) {
	var once Once
	f := func() {}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(f)
		}
	})
}

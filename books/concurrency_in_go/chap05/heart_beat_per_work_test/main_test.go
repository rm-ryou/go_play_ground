package main

import (
	"testing"
)

// func Test_doWork_GenerateAllNumbers(t *testing.T) {
// 	done := make(chan any)
// 	defer close(done)

// 	intSlice := []int{0, 1, 2, 3, 5}
// 	_, results := doWork(done, intSlice...)

// 	for i, expected := range intSlice {
// 		select {
// 		case r := <-results:
// 			if r != expected {
// 				t.Errorf(
// 					"index %v: expected %v, but received %v.",
// 					i,
// 					expected,
// 					r,
// 				)
// 			}
// 		case <-time.After(1 * time.Second):
// 			t.Fatal("test time out")
// 		}
// 	}
// }

func Test_doWork_GenerateAllNumbers_with_heartbeat(t *testing.T) {
	done := make(chan any)
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := doWork(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
		}
		i++
	}
}

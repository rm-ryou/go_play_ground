package main

import "testing"

func Test_fn(t *testing.T) {
	N := 4
	W := 14
	arry := []int{3, 2, 6, 5}
	expected := true

	actual := fn(N, W, arry)
	if actual != expected {
		t.Errorf("expected is %v. but actual is %v", expected, actual)
	}
}

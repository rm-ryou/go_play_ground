package main

import "testing"

func Test_multiple(t *testing.T) {
	n := 3
	a := []int{3, 2, 9}
	b := []int{5, 7, 4}
	expected := 7

	actual := multiple(a, b, n)
	if actual != expected {
		t.Errorf("expected is %d. But actual is %d", expected, actual)
	}
}

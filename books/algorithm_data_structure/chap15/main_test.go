package main

import "testing"

func Test_Kruskal(t *testing.T) {
	N := 8
	M := 12
	a := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 4, 6}
	b := []int{3, 5, 7, 3, 4, 6, 4, 5, 7, 7, 6, 7}
	w := []int{5, 6, 3, 8, 4, 3, 9, 10, 5, 6, 2, 7}

	want := 31

	t.Run("全域木の最小値を取得する", func(t *testing.T) {
		act := Kruskal(N, M, a, b, w)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

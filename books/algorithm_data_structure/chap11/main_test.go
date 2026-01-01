package main

import "testing"

func TestSizeOfGroup(t *testing.T) {
	N := 13
	M := 14

	// 1, 2
	// 1, 3
	// 1, 4
	// 2, 5
	// 2, 6
	// 3, 6
	// 4, 7
	// 5, 6
	// 5, 8
	// 6, 7
	// 7, 9
	// 8, 9
	// 11, 13
	// 12, 13
	a := []int{1, 1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 11, 12}
	b := []int{2, 3, 4, 5, 6, 6, 7, 6, 8, 7, 9, 9, 13, 13}

	t.Run("グラフの連結成分の個数を取得する", func(t *testing.T) {
		want := 3
		act := SizeOfGroup(N, M, a, b)

		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

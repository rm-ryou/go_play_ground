package main

import "testing"

func Test_Flog(t *testing.T) {
	t.Run("N-1の足場に着くまでに要する最小のコストを返す", func(t *testing.T) {
		N := 7
		h := []int{2, 9, 4, 5, 1, 6, 10}
		want := 8

		act := Flog(N, h)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_FlogPushBased(t *testing.T) {
	t.Run("N-1の足場に着くまでに要する最小のコストを返す", func(t *testing.T) {
		N := 7
		h := []int{2, 9, 4, 5, 1, 6, 10}
		want := 8

		act := FlogPushBased(N, h)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_Knapsack(t *testing.T) {
	t.Run("wを超えないように価値の総和を最大化する", func(t *testing.T) {
		N := 6
		w := []int{2, 1, 3, 2, 1, 5}
		v := []int{3, 2, 6, 1, 3, 85}
		W := 9
		want := 94

		act := Knapsack(N, w, v, W)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_EditDistance(t *testing.T) {
	t.Run("文字列SをTに変換する最小のコスト", func(t *testing.T) {
		S := "logistic"
		T := "algorithm"
		want := 6

		act := EditDistance(S, T)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

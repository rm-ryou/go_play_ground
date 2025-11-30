package main

import "testing"

func Test_SumRecv(t *testing.T) {
	t.Run("1からnまでの和を返すこと", func(t *testing.T) {
		want := 55
		n := 10

		act := SumRecv(n)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_GCD(t *testing.T) {
	t.Run("最大公約数を求めること", func(t *testing.T) {
		want := 3

		act := GCD(51, 15)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_FiboRecv(t *testing.T) {
	t.Run("フィボナッチ数を返すこと", func(t *testing.T) {
		want := 13

		act := FiboRecv(7)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_FiboMemo(t *testing.T) {
	t.Run("フィボナッチ数を返すこと", func(t *testing.T) {
		want := 7778742049

		memo := make([]int, 50)
		for i := range memo {
			memo[i] = -1
		}

		FiboMemo(memo, 49)
		act := memo[49]
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

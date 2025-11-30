package main

import "testing"

func Test_BasicLinearSearch(t *testing.T) {
	t.Run("ary内にkeyが存在する時、trueを返すこと", func(t *testing.T) {
		want := true
		ary := []int{1, 2, 3, 4}
		key := 2

		act := BasicLinearSearch(ary, key)
		if act != want {
			t.Errorf("want: %t, got: %t", want, act)
		}
	})

	t.Run("ary内にkeyが存在しない時、falseを返すこと", func(t *testing.T) {
		want := false
		ary := []int{1, 2, 3, 4}
		key := 5

		act := BasicLinearSearch(ary, key)
		if act != want {
			t.Errorf("want: %t, got: %t", want, act)
		}
	})
}

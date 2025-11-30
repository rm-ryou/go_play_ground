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

func Test_FindIndex(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		key  int
		want int
	}{
		{

			name: "ary内にkeyが存在する時、添字を返す",
			ary:  []int{1, 2, 3, 4},
			key:  2,
			want: 1,
		},
		{

			name: "ary内にkeyが存在する時、-1を返す",
			ary:  []int{1, 2, 3, 4},
			key:  5,
			want: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := FindIndex(tc.ary, tc.key)

			if act != tc.want {
				t.Errorf("want: %d, got: %d", tc.want, act)
			}
		})
	}
}

func Test_FindMinValue(t *testing.T) {
	t.Run("ary内の最小値を取得すること", func(t *testing.T) {
		want := 1
		ary := []int{1, 2, 3, 4}

		act := FindMinValue(ary)
		if act != want {
			t.Errorf("want: %d, got: %d", want, act)
		}
	})
}

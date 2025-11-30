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

func Test_IsCreateValueFromAry(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		w    int
		want bool
	}{
		{
			name: "aryの要素からwを作成できる時、trueを返す",
			ary:  []int{1, 2, 3, 5, 11},
			w:    10,
			want: true,
		},
		{
			name: "aryの要素からwを作成できない時、falseを返す",
			ary:  []int{1, 5, 8, 11},
			w:    10,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsCreateValueFromAry(len(tc.ary), tc.w, tc.ary)

			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

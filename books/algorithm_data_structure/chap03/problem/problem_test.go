package problem

import "testing"

func Test_CountValue(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		key  int
		want int
	}{
		{
			name: "aryのなかにkeyがある時、keyの個数を返す",
			ary:  []int{1, 1, 2, 3},
			key:  1,
			want: 2,
		},
		{
			name: "aryのなかにkeyがない時、0を返す",
			ary:  []int{1, 1, 2, 3},
			key:  9,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := CountValue(tc.ary, tc.key)
			if act != tc.want {
				t.Errorf("want: %d, got: %d", tc.want, act)
			}
		})
	}
}

func Test_FetchSecondMinValue(t *testing.T) {
	t.Run("ary内の2番目に小さい値を取得すること", func(t *testing.T) {
		want := 1
		ary := []int{1, 1, 2, 3, 4}

		act := FetchSecondMinValue(ary)
		if act != want {
			t.Errorf("want: %d, got: %d", want, act)
		}
	})
}

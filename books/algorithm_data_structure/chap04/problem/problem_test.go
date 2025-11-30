package problem

import "testing"

func Test_Tribo(t *testing.T) {
	t.Run("トリボナッチ数列を返す", func(t *testing.T) {
		want := 13

		act := Tribo(7)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_TriboMemo(t *testing.T) {
	t.Run("トリボナッチ数列を返す", func(t *testing.T) {
		want := 7046319384

		idx := 40
		memo := make([]int, idx+1)
		for i, _ := range memo {
			memo[i] = -1
		}

		TriboMemo(memo, idx)
		act := memo[idx]
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_Problem755(t *testing.T) {
	testCases := []struct {
		name string
		k    int
		want int
	}{
		{
			name: "return 4",
			k:    575,
			want: 4,
		},
		{
			name: "return 13",
			k:    3600,
			want: 13,
		},
		{
			name: "return 26484",
			k:    999999999,
			want: 26484,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := 0
			Problem755(tc.k, 0, 0, &act)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

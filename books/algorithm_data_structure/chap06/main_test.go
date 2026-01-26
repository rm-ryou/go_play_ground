package main

import "testing"

func Test_CalcMostMinPair(t *testing.T) {
	t.Run("a, bから、K以上で最小値のペアの合計を取得する", func(t *testing.T) {
		n := 3
		k := 10
		a := []int{8, 5, 4}
		b := []int{4, 1, 9}
		want := 12

		act := CalcMostMinPair(n, k, a, b)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_Shooting(t *testing.T) {
	type testCase struct {
		name string
		N    int
		H, S []int
		want int
	}

	testCases := []testCase{
		{
			name: "return 23",
			N:    4,
			H:    []int{5, 12, 14, 21},
			S:    []int{6, 4, 7, 2},
			want: 23,
		},
		{
			name: "return 105",
			N:    6,
			H:    []int{100, 100, 100, 100, 100, 1},
			S:    []int{1, 1, 1, 1, 1, 30},
			want: 105,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Shooting(tc.N, tc.H, tc.S)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

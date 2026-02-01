package problem

import "testing"

func Test_ComplessMap(t *testing.T) {
	t.Run("座標圧縮を行う", func(t *testing.T) {
		n := 5
		a := []int{12, 43, 7, 15, 9}
		want := []int{2, 4, 0, 3, 1}

		act := ComplessMap(n, a)
		if len(a) != len(want) {
			t.Errorf("invalid ary: want len: %d, act len: %d", len(want), len(act))
		}

		for i := range n {
			if act[i] != want[i] {
				t.Errorf("invalid ary: idx: %d, want: %d, act: %d", i, want[i], act[i])
			}
		}
	})
}

func Test_Sunke(t *testing.T) {
	type testCase struct {
		name    string
		N       int
		A, B, C []int
		want    int
	}

	testCases := []testCase{
		{
			name: "return 3",
			N:    2,
			A:    []int{1, 5},
			B:    []int{2, 4},
			C:    []int{3, 6},
			want: 3,
		},
		{
			name: "return 27",
			N:    3,
			A:    []int{1, 1, 1},
			B:    []int{2, 2, 2},
			C:    []int{3, 3, 3},
			want: 27,
		},
		{
			name: "return 87",
			N:    6,
			A:    []int{3, 14, 159, 2, 6, 53},
			B:    []int{58, 9, 79, 323, 84, 6},
			C:    []int{2643, 383, 2, 79, 50, 288},
			want: 87,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Sunke(tc.N, tc.A, tc.B, tc.C)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_Darts(t *testing.T) {
	type testCase struct {
		name string
		N, M int
		a    []int
		want int
	}

	testCases := []testCase{
		{
			name: "aの要素から重複ありで4つ選び、Mを超えない最大値を返す",
			N:    4,
			M:    50,
			a:    []int{3, 14, 15, 9},
			want: 48,
		},
		{
			name: "aの要素から重複ありで4つ選び、Mを超えない最大値を返す",
			N:    3,
			M:    21,
			a:    []int{16, 11, 2},
			want: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Darts(tc.N, tc.M, tc.a)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_AggressiveCows(t *testing.T) {
	t.Run("M個選んだ要素の内、2つの距離の最小値の最大値", func(t *testing.T) {
		n := 5
		m := 3
		a := []int{1, 2, 8, 4, 9}
		want := 3

		act := AggressiveCows(n, m, a)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

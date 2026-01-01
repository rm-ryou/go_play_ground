package problem

import "testing"

func TestBridge(t *testing.T) {
	type testCase struct {
		name string
		N, M int
		A, B []int
		want int
	}

	testCases := []testCase{
		{
			name: "連結でなくなる変の数を返す",
			N:    7,
			M:    7,
			A:    []int{1, 2, 3, 4, 4, 5, 6},
			B:    []int{3, 7, 4, 5, 6, 6, 7},
			want: 4,
		},
		{
			name: "連結でなくなる辺がない時、0を返す",
			N:    3,
			M:    3,
			A:    []int{1, 1, 2},
			B:    []int{2, 3, 3},
			want: 0,
		},
		{
			name: "連結でなくなる辺が全ての時、総数を返す",
			N:    6,
			M:    5,
			A:    []int{1, 2, 3, 4, 5},
			B:    []int{2, 3, 4, 5, 6},
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Bridge(tc.N, tc.M, tc.A, tc.B)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func TestDecayed(t *testing.T) {
	type testCase struct {
		name string
		N, M int
		A, B []int
		want []int
	}

	testCases := []testCase{
		{
			name: "return [0, 0, 4, 5, 6]",
			N:    4,
			M:    5,
			A:    []int{1, 3, 1, 2, 1},
			B:    []int{2, 4, 3, 3, 4},
			want: []int{0, 0, 4, 5, 6},
		},
		{
			name: "return [8, 9, 12, 14, 15]",
			N:    6,
			M:    5,
			A:    []int{2, 1, 5, 3, 4},
			B:    []int{3, 2, 6, 4, 5},
			want: []int{8, 9, 12, 14, 15},
		},
		{
			name: "return [1]",
			N:    2,
			M:    1,
			A:    []int{1},
			B:    []int{2},
			want: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Decayed(tc.N, tc.M, tc.A, tc.B)

			if len(act) != tc.M {
				t.Error("no matched size")
			}

			for i := range tc.M {
				if act[i] != tc.want[i] {
					t.Errorf("want: %d, act: %d", tc.want[i], act[i])
				}
			}
		})
	}
}

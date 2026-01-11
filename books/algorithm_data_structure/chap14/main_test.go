package main

import "testing"

func Test_BellmanFord(t *testing.T) {
	type testCase struct {
		name       string
		N, M, s, v int
		a, b, w    []int
		want       int
	}

	testCases := []testCase{
		{
			name: "最短経路を取得する",
			N:    6,
			M:    12,
			s:    0,
			v:    5,
			a:    []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 4, 4, 4},
			b:    []int{1, 3, 2, 3, 4, 3, 4, 5, 1, 2, 3, 5},
			w:    []int{3, 100, 50, 57, -4, -10, -5, 100, -5, 57, 25, 8},
			want: 7,
		},
		{
			name: "負閉路がある時、-1を返す",
			N:    6,
			M:    12,
			s:    0,
			v:    5,
			a:    []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 4, 4, 4},
			b:    []int{1, 3, 2, 3, 4, 3, 4, 5, 1, 2, 3, 5},
			w:    []int{3, 100, -50, 57, -4, -10, -5, 100, -5, 57, -25, 8},
			want: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := BellmanFord(tc.N, tc.M, tc.s, tc.v, tc.a, tc.b, tc.w)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}

}

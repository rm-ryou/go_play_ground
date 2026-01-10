package main

import "testing"

func Test_IsSTPathExists(t *testing.T) {
	type testCase struct {
		name       string
		n, m, s, t int
		a, b       []int
		want       bool
	}

	testCases := []testCase{
		{
			name: "s-tパスが存在する時、trueを返す",
			n:    8,
			m:    12,
			s:    2,
			t:    7,
			a:    []int{0, 0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 6},
			b:    []int{1, 2, 3, 2, 4, 5, 5, 7, 5, 6, 6, 7},
			want: true,
		},
		{
			name: "s-tパスが存在しない時、trueを返す",
			n:    8,
			m:    12,
			s:    2,
			t:    3,
			a:    []int{0, 0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 6},
			b:    []int{1, 2, 3, 2, 4, 5, 5, 7, 5, 6, 6, 7},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsSTPathExists(tc.n, tc.m, tc.s, tc.t, tc.a, tc.b)

			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

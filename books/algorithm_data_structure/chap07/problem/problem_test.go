package problem

import "testing"

func Test_Pair(t *testing.T) {
	type testCase struct {
		name string
		a, b []int
		want int
	}

	testCases := []testCase{
		{
			name: "return 3",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: 3,
		},
		{
			name: "return 0",
			a:    []int{6, 5, 7},
			b:    []int{2, 3, 3},
			want: 0,
		},
		{
			name: "return 3",
			a:    []int{8, 1, 2, 7, 5},
			b:    []int{7, 6, 3, 1, 4},
			want: 3,
		},
	}

	t.Run("an < bnなるペアが最大何個作れるか", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				act := Pair(tc.a, tc.b)
				if act != tc.want {
					t.Errorf("want: %d, act: %d", tc.want, act)
				}
			})
		}
	})
}

func Test_PlainPoints(t *testing.T) {
	type testCase struct {
		name string
		n    int
		a, b []Cod
		want int
	}

	testCases := []testCase{
		{
			name: "return 2",
			n:    3,
			a: []Cod{
				{X: 2, Y: 0},
				{X: 3, Y: 1},
				{X: 1, Y: 3},
			},
			b: []Cod{
				{X: 4, Y: 2},
				{X: 0, Y: 4},
				{X: 5, Y: 5},
			},
			want: 2,
		},
		{
			name: "return 2",
			n:    3,
			a: []Cod{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
				{X: 5, Y: 2},
			},
			b: []Cod{
				{X: 2, Y: 3},
				{X: 3, Y: 4},
				{X: 4, Y: 5},
			},
			want: 2,
		},
		{
			name: "return 0",
			n:    2,
			a: []Cod{
				{X: 2, Y: 2},
				{X: 3, Y: 3},
			},
			b: []Cod{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			},
			want: 0,
		},
		{
			name: "return 5",
			n:    5,
			a: []Cod{
				{X: 0, Y: 0},
				{X: 7, Y: 3},
				{X: 2, Y: 2},
				{X: 4, Y: 8},
				{X: 1, Y: 6},
			},
			b: []Cod{
				{X: 8, Y: 5},
				{X: 6, Y: 9},
				{X: 5, Y: 4},
				{X: 9, Y: 1},
				{X: 3, Y: 7},
			},
			want: 5,
		},
	}

	t.Run("ax<bx && ay<byなるペアの最大数を返す", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				act := PlainPoints(tc.n, tc.a, tc.b)
				if act != tc.want {
					t.Errorf("want: %d, act: %d", tc.want, act)
				}
			})
		}
	})

}

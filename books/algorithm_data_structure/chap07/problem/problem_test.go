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

func Test_Megalo(t *testing.T) {
	type testCase struct {
		name string
		n    int
		a, b []int
		want bool
	}

	testCases := []testCase{
		{
			name: "return true",
			n:    5,
			a:    []int{2, 1, 1, 4, 3},
			b:    []int{4, 9, 8, 9, 12},
			want: true,
		},
		{
			name: "return false",
			n:    3,
			a:    []int{334, 334, 334},
			b:    []int{1000, 1000, 1000},
			want: false,
		},
		{
			name: "return true",
			n:    30,
			a: []int{
				384,
				1725,
				170,
				4,
				2,
				578,
				702,
				143,
				1420,
				24,
				849,
				76,
				85,
				444,
				719,
				470,
				1137,
				455,
				110,
				15,
				368,
				104,
				3,
				366,
				7,
				610,
				152,
				4,
				292,
				334,
			},
			b: []int{
				8895,
				9791,
				1024,
				11105,
				6,
				1815,
				3352,
				5141,
				6980,
				1602,
				999,
				7586,
				5570,
				4991,
				11090,
				10708,
				4547,
				9003,
				9901,
				8578,
				3692,
				1286,
				4,
				12143,
				6649,
				2374,
				7324,
				7042,
				11386,
				5720,
			},
			want: true,
		},
	}

	t.Run("aiの時間を要し、biの時刻が締め切りの中、全てのタスクをこなせられるか", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				act := Megalo(tc.n, tc.a, tc.b)
				if act != tc.want {
					t.Errorf("want: %t, act: %t", tc.want, act)
				}
			})
		}
	})
}

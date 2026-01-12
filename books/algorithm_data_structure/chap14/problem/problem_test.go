package problem

import "testing"

func Test_LongestPath(t *testing.T) {
	type testCase struct {
		name string
		N, M int
		x, y []int
		want int
	}

	testCases := []testCase{
		{
			name: "有効パスの最長値を返す",
			N:    4,
			M:    5,
			x:    []int{1, 1, 3, 2, 3},
			y:    []int{2, 3, 2, 4, 4},
			want: 3,
		},
		{
			name: "return 2",
			N:    6,
			M:    3,
			x:    []int{2, 4, 5},
			y:    []int{3, 5, 6},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := LongestPath(tc.N, tc.M, tc.x, tc.y)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_ScoreAttack(t *testing.T) {
	type testCase struct {
		name    string
		N, M    int
		x, y, w []int
		want    int
	}

	testCases := []testCase{
		{
			name: "最大のスコアを返す",
			N:    3,
			M:    3,
			x:    []int{1, 2, 1},
			y:    []int{2, 3, 3},
			w:    []int{4, 3, 5},
			want: 7,
		},
		{
			name: "最大のスコアを返す",
			N:    6,
			M:    5,
			x:    []int{1, 2, 3, 4, 5},
			y:    []int{2, 3, 4, 5, 6},
			w: []int{
				-1000000000,
				-1000000000,
				-1000000000,
				-1000000000,
				-1000000000,
			},
			want: -5000000000,
		},
		{
			name: "サイクルがあり、スコアが無限に増やせる時、-1を返す",
			N:    2,
			M:    2,
			x:    []int{1, 2},
			y:    []int{2, 1},
			w:    []int{1, 1},
			want: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := ScoreAttack(tc.N, tc.M, tc.x, tc.y, tc.w)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

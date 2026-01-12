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

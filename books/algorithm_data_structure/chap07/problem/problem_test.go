package problem

import "testing"

func Test_Pair(t *testing.T) {
	type testCase struct {
		name string
		a, b []int
		want int
	}

	testCases := []testCase{
		// {
		// 	name: "return 3",
		// 	a:    []int{1, 2, 3},
		// 	b:    []int{4, 5, 6},
		// 	want: 3,
		// },
		// {
		// 	name: "return 0",
		// 	a:    []int{6, 5, 7},
		// 	b:    []int{2, 3, 3},
		// 	want: 0,
		// },
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

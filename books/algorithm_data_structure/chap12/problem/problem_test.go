package problem

import "testing"

func Test_Energy(t *testing.T) {
	type testCase struct {
		name  string
		n, m  int
		price []int
		count []int
		want  int
	}

	testCases := []testCase{
		{
			name:  "return 12",
			n:     2,
			m:     5,
			price: []int{4, 2},
			count: []int{9, 4},
			want:  12,
		},
		{
			name:  "return 130",
			n:     4,
			m:     30,
			price: []int{6, 2, 3, 7},
			count: []int{18, 5, 10, 9},
			want:  130,
		},
		{
			name:  "return 100000000000000",
			n:     1,
			m:     100000,
			price: []int{1000000000},
			count: []int{100000},
			want:  100000000000000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Energy(tc.n, tc.m, tc.price, tc.count)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

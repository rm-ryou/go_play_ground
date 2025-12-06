package problem

import "testing"

func Test_Vacation(t *testing.T) {
	testCases := []struct {
		name    string
		N       int
		a, b, c []int
		want    int
	}{
		{
			name: "return 210",
			N:    3,
			a:    []int{10, 20, 30},
			b:    []int{40, 50, 60},
			c:    []int{70, 80, 90},
			want: 210,
		},
		{
			name: "return 100",
			N:    1,
			a:    []int{100},
			b:    []int{10},
			c:    []int{1},
			want: 100,
		},
		{
			name: "return 46",
			N:    7,
			a:    []int{6, 8, 2, 7, 4, 2, 7},
			b:    []int{7, 8, 5, 8, 6, 3, 5},
			c:    []int{8, 3, 2, 6, 8, 4, 1},
			want: 46,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Vacation(tc.N, tc.a, tc.b, tc.c)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

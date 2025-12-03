package main

import "testing"

func Test_Flog(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		ary  []int
		want int
	}{
		{
			name: "return 30",
			n:    4,
			ary:  []int{10, 30, 40, 20},
			want: 30,
		},
		{
			name: "return 0",
			n:    2,
			ary:  []int{10, 10},
			want: 0,
		},
		{
			name: "return 40",
			n:    6,
			ary:  []int{30, 10, 60, 10, 60, 50},
			want: 40,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Flog(tc.n, tc.ary)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

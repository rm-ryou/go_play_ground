package problem

import "testing"

func Test_Vacation(t *testing.T) {
	type testCase struct {
		name    string
		n       int
		a, b, c []int
		want    int
	}

	testCases := []testCase{
		{
			name: "return 210",
			n:    3,
			a:    []int{10, 20, 30},
			b:    []int{40, 50, 60},
			c:    []int{70, 80, 90},
			want: 210,
		},
		{
			name: "return 46",
			n:    7,
			a:    []int{6, 8, 2, 7, 4, 2, 7},
			b:    []int{7, 8, 5, 8, 6, 3, 5},
			c:    []int{8, 3, 2, 6, 8, 4, 1},
			want: 46,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Vacation(tc.n, tc.a, tc.b, tc.c)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_SumOfPart(t *testing.T) {
	type testCase struct {
		name string
		n, w int
		a    []int
		want bool
	}

	testCases := []testCase{
		{
			name: "aのいくつかの要素をたし、wを作ることができる時、trueを返すこと",
			n:    5,
			w:    10,
			a:    []int{1, 2, 3, 5, 11},
			want: true,
		},
		{
			name: "aのいくつかの要素をたし、wを作ることができない時、falseを返すこと",
			n:    5,
			w:    4,
			a:    []int{2, 6, 3, 5, 11},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := SumOfPart(tc.n, tc.w, tc.a)
			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

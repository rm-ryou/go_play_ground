package problem

import "testing"

func Test_Compression(t *testing.T) {
	ary := []int{12, 43, 7, 15, 9}
	want := []int{2, 4, 0, 3, 1}

	act := Compression(ary)
	for i := range len(want) {
		if act[i] != want[i] {
			t.Errorf("want: %v\nact: %v", want, act)
		}
	}
}

func Test_Snuke(t *testing.T) {
	testCases := []struct {
		name    string
		a, b, c []int
		want    int
	}{
		{
			name: "return 3",
			a:    []int{1, 5},
			b:    []int{2, 4},
			c:    []int{3, 6},
			want: 3,
		},
		{
			name: "return 27",
			a:    []int{1, 1, 1},
			b:    []int{2, 2, 2},
			c:    []int{3, 3, 3},
			want: 27,
		},
		{
			name: "return 87",
			a:    []int{3, 14, 159, 2, 6, 53},
			b:    []int{58, 9, 79, 323, 84, 6},
			c:    []int{2643, 383, 2, 79, 50, 288},
			want: 87,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Snuke(tc.a, tc.b, tc.c)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_Dart(t *testing.T) {
	testCases := []struct {
		name string
		m    int
		a    []int
		want int
	}{
		{
			name: "return 20",
			m:    21,
			a:    []int{16, 11, 2},
			want: 20,
		},
		{
			name: "return 48",
			m:    50,
			a:    []int{3, 14, 15, 9},
			want: 48,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Dart(tc.m, tc.a)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

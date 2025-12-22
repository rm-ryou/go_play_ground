package main

import "testing"

func Test_Coin(t *testing.T) {
	testCases := []struct {
		name  string
		value int
		nums  []int
		want  int
	}{
		{
			name:  "return 5",
			value: 22,
			nums:  []int{3, 1, 4, 2},
			want:  5,
		},
		{
			name:  "return 11",
			value: 314,
			nums:  []int{10, 10, 10, 10},
			want:  11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Coin(tc.value, tc.nums)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_Schedule(t *testing.T) {
	N := 3
	starts := []int{1, 4, 6}
	ends := []int{5, 6, 8}
	want := 2

	t.Run("実行可能なタスクの最大値を返すこと", func(t *testing.T) {
		act := Schedule(N, starts, ends)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_MultipleArray(t *testing.T) {
	testCases := []struct {
		name    string
		n       int
		buttons []int
		nums    []int
		want    int
	}{
		{
			name:    "return 7",
			n:       3,
			buttons: []int{3, 2, 9},
			nums:    []int{5, 7, 4},
			want:    7,
		},
		{
			name:    "return 22",
			n:       7,
			buttons: []int{3, 4, 5, 2, 5, 5, 9},
			nums:    []int{1, 1, 9, 6, 3, 8, 7},
			want:    22,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := MultipleArray(tc.n, tc.buttons, tc.nums)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

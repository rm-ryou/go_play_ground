package main

import "testing"

func Test_BinarySearch(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		key  int
		want bool
	}{
		{
			name: "aryにkeyが含まれている時、trueを返す",
			ary:  []int{3, 5, 8, 10, 14, 17, 21, 39},
			key:  17,
			want: true,
		},
		{
			name: "aryにkeyが含まれていない時、falseを返す",
			ary:  []int{3, 5, 8, 10, 14, 17, 21, 39},
			key:  9,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := BinarySearch(tc.ary, tc.key)
			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

func Test_IsGuessNumberUnderN(t *testing.T) {
	testCases := []struct {
		name     string
		min, max int
		limit    int
		want     bool
	}{
		{
			name:  "limit回までに数値を当てることができる場合、trueを返す",
			min:   20,
			max:   35,
			limit: 4,
			want:  true,
		},
		{
			name:  "limit回までに数値を当てることができない場合、falseを返す",
			min:   0,
			max:   100,
			limit: 3,
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsGuessNumberUnderN(tc.min, tc.max, tc.limit)
			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

func Test_MinSumOverK(t *testing.T) {
	a := []int{8, 5, 4}
	b := []int{4, 1, 9}
	k := 10
	want := 12

	act := MinSumOverK(a, b, k)
	if act != want {
		t.Errorf("want: %d, act: %d", want, act)
	}
}

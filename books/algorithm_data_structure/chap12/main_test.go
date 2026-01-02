package main

import (
	"testing"
)

func assertSlice(t *testing.T, want, act []int) {
	t.Helper()

	if len(want) != len(act) {
		t.Errorf("different size. want: %d, act: %d", len(want), len(act))
	}

	for i := range len(want) {
		if want[i] != act[i] {
			t.Errorf("%d is invalid. want: %d, act: %d", i, want[i], act[i])
		}
	}
}

func Test_InsertionSort(t *testing.T) {
	type testCase struct {
		name string
		data []int
		want []int
	}

	testCases := []testCase{
		{
			name: "昇順にソートすること",
			data: []int{5, 2, 4, 3, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "同一値が含まれている時、順序関係なくソートすること",
			data: []int{1, 1, 2, 1, 1},
			want: []int{1, 1, 1, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := InsertionSort(tc.data)

			assertSlice(t, tc.want, act)
		})
	}
}

func Test_MergeSort(t *testing.T) {
	type testCase struct {
		name string
		data []int
		want []int
	}

	testCases := []testCase{
		{
			name: "昇順にソートすること",
			data: []int{5, 2, 4, 3, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "同一値が含まれている時、順序関係なくソートすること",
			data: []int{1, 1, 2, 1, 1},
			want: []int{1, 1, 1, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := MergeSort(tc.data)

			assertSlice(t, tc.want, act)
		})
	}
}

func Test_QuickSort(t *testing.T) {
	type testCase struct {
		name string
		data []int
		want []int
	}

	testCases := []testCase{
		{
			name: "昇順にソートすること",
			data: []int{5, 2, 4, 3, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		// {
		// 	name: "同一値が含まれている時、順序関係なくソートすること",
		// 	data: []int{1, 1, 2, 1, 1},
		// 	want: []int{1, 1, 1, 1, 2},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := QuickSort(tc.data)

			assertSlice(t, tc.want, act)
		})
	}
}

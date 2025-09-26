package main

import "testing"

func Test_partOfSumUnderK(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		w        int
		k        int
		arry     []int
		expected bool
	}{
		{
			name:     "k以内でwを作れる時はtrueを返す",
			n:        5,
			w:        10,
			k:        3,
			arry:     []int{1, 2, 4, 5, 11},
			expected: true,
		},
		{
			name:     "k以内でwを作れない時はfalseを返す",
			n:        5,
			w:        10,
			k:        2,
			arry:     []int{1, 2, 4, 5, 11},
			expected: false,
		},
		{
			name:     "wが作れない時は falseを返す",
			n:        5,
			w:        10,
			k:        5,
			arry:     []int{1, 2, 6, 5, 11},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := partOfSumUnderK(tc.n, tc.w, tc.k, tc.arry)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_partOfSumCounter(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		w        int
		arry     []int
		expected int
	}{
		{
			name:     "return 7",
			n:        3,
			w:        10,
			arry:     []int{2, 3, 5},
			expected: 6,
		},
		{
			name:     "return 10",
			n:        10,
			w:        10,
			arry:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := partOfSumCounter(tc.n, tc.w, tc.arry)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_partOfSumBasic(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		w        int
		arry     []int
		expected bool
	}{
		{
			name:     "return true",
			n:        5,
			w:        10,
			arry:     []int{1, 2, 4, 5, 11},
			expected: true,
		},
		{
			name:     "return true",
			n:        5,
			w:        7,
			arry:     []int{1, 2, 4, 5, 11},
			expected: true,
		},
		{
			name:     "return false",
			n:        5,
			w:        10,
			arry:     []int{1, 2, 6, 5, 11},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := partOfSumBasic(tc.n, tc.w, tc.arry)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_vacation(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		arryA    []int
		arryB    []int
		arryC    []int
		expected int
	}{
		{
			name:     "return 210",
			n:        3,
			arryA:    []int{10, 20, 30},
			arryB:    []int{40, 50, 60},
			arryC:    []int{70, 80, 90},
			expected: 210,
		},
		{
			name:     "return 100",
			n:        1,
			arryA:    []int{100},
			arryB:    []int{10},
			arryC:    []int{1},
			expected: 100,
		},
		{
			name:     "return 46",
			n:        7,
			arryA:    []int{6, 8, 2, 7, 4, 2, 7},
			arryB:    []int{7, 8, 5, 8, 6, 3, 5},
			arryC:    []int{8, 3, 2, 6, 8, 4, 1},
			expected: 46,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := vacation(tc.n, tc.arryA, tc.arryB, tc.arryC)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

// 7
// 6 7 8
// 8 8 3
// 2 5 2
// 7 8 6
// 4 6 8
// 2 3 4
// 7 5 1

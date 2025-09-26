package main

import "testing"

func Test_megalo(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		d, t     []int
		expected bool
	}{
		{
			name:     "return true",
			n:        5,
			d:        []int{2, 1, 1, 4, 3},
			t:        []int{4, 9, 8, 9, 12},
			expected: true,
		},
		{
			name:     "return false",
			n:        3,
			d:        []int{334, 334, 334},
			t:        []int{1000, 1000, 1000},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := megalo(tc.n, tc.d, tc.t)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_plan2d(t *testing.T) {
	testCases := []struct {
		name      string
		n         int
		red, blue [][]int
		expected  int
	}{
		{
			name: "return 2",
			n:    3,
			red: [][]int{
				{2, 0},
				{3, 1},
				{1, 3},
			},
			blue: [][]int{
				{4, 2},
				{0, 4},
				{5, 5},
			},
			expected: 2,
		},
		{
			name: "return 0",
			n:    2,
			red: [][]int{
				{2, 2},
				{3, 3},
			},
			blue: [][]int{
				{0, 0},
				{1, 1},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := plan2d(tc.n, tc.red, tc.blue)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_prob71(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		a, b     []int
		expected int
	}{
		{
			name:     "return true",
			n:        5,
			a:        []int{2, 1, 1, 4, 3},
			b:        []int{4, 9, 8, 9, 12},
			expected: 5,
		},
		{
			name:     "return false",
			n:        3,
			a:        []int{9, 1, 3, 5, 11},
			b:        []int{1, 2, 8, 6, 100},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := prob71(tc.n, tc.a, tc.b)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

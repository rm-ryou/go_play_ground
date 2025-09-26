package main

import "testing"

func Test_canCreateW(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		w        int
		expected bool
	}{
		{
			name:     "return true",
			arry:     []int{1, 2, 4, 5, 10, 11},
			w:        10,
			expected: true,
		},
		{
			name:     "return false",
			arry:     []int{1, 5, 8, 11},
			w:        10,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := canCreateW(tc.arry, tc.w)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

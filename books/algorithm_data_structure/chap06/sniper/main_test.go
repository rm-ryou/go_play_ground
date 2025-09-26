package main

import "testing"

func Test_sniper(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		h        []int
		s        []int
		expected int
	}{
		{
			name:     "return 23",
			n:        4,
			h:        []int{5, 12, 14, 21},
			s:        []int{6, 7, 4, 2},
			expected: 23,
		},
		{
			name:     "return 105",
			n:        6,
			h:        []int{100, 100, 100, 100, 100, 1},
			s:        []int{1, 1, 1, 1, 1, 30},
			expected: 105,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sniper(tc.n, tc.h, tc.s)

			if actual != tc.expected {
				t.Errorf("expected is %d. but actual is %d", tc.expected, actual)
			}
		})
	}
}

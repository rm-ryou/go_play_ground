package main

import "testing"

func Test_findMinValue(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		expected int
	}{
		{
			name:     "return min value",
			arry:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMinValue(tc.arry)

			if actual != tc.expected {
				t.Errorf("expected is %d. but actual is %v", tc.expected, actual)
			}
		})
	}
}

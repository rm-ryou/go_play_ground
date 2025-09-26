package main

import "testing"

func Test_linearSearch(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		val      int
		expected bool
	}{
		{
			name:     "return true",
			arry:     []int{1, 2, 3, 4, 5},
			val:      2,
			expected: true,
		},
		{
			name:     "return false",
			arry:     []int{1, 2, 3, 4, 5},
			val:      10,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := linearSearch(tc.arry, tc.val)

			if actual != tc.expected {
				t.Errorf("expected is %v but actual %v", tc.expected, actual)
			}
		})
	}
}

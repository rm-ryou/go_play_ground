package main

import "testing"

func Test_linearSearchWithIndex(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		val      int
		expected int
	}{
		{
			name:     "return 0",
			arry:     []int{1, 2, 3, 4, 5},
			val:      1,
			expected: 0,
		},
		{
			name:     "when val not exists in arry. return -1",
			arry:     []int{1, 2, 3, 4, 5},
			val:      7,
			expected: -1,
		},
		{
			name:     "when find multi values. return last found index",
			arry:     []int{1, 2, 3, 4, 5, 1},
			val:      1,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := linearSearchWithIndex(tc.arry, tc.val)

			if actual != tc.expected {
				t.Errorf("expected is %v. but actual is %v", tc.expected, actual)
			}
		})
	}
}

package main

import "testing"

func Test_tribonacci(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "return 0",
			n:        1,
			expected: 0,
		},
		{
			name:     "return 7",
			n:        6,
			expected: 7,
		},
		{
			name:     "return 44",
			n:        9,
			expected: 44,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tribonacci(tc.n)

			if actual != tc.expected {
				t.Errorf("expected is %d. but actual is %d", tc.expected, actual)
			}
		})
	}
}

func Test_tribonacciMemo(t *testing.T) {
	testCases := []struct {
		name     string
		setupFn  func(n int)
		n        int
		expected int
	}{
		{
			name: "return 0",
			setupFn: func(n int) {
				memo = make([]int, n+1)
			},
			n:        1,
			expected: 0,
		},
		{
			name: "return 7",
			setupFn: func(n int) {
				memo = make([]int, n+1)
			},
			n:        6,
			expected: 7,
		},
		{
			name: "return 44",
			setupFn: func(n int) {
				memo = make([]int, n+1)
			},
			n:        9,
			expected: 44,
		},
		{
			name: "return 81",
			setupFn: func(n int) {
				memo = make([]int, n+1)
			},
			n:        10,
			expected: 81,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupFn(tc.n)
			actual := tribonacciMemo(tc.n)

			if actual != tc.expected {
				t.Errorf("expected is %d. but actual is %d", tc.expected, actual)
			}
		})
	}
}

func Test_fn755(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "return 4",
			n:        575,
			expected: 4,
		},
		{
			name:     "return 13",
			n:        3600,
			expected: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := 0
			fn755(tc.n, 0, 0, &actual)

			if actual != tc.expected {
				t.Errorf("expected is %d. but actual is %d", tc.expected, actual)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_darts(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		m        int
		expected int
	}{
		{
			name:     "fooo",
			arry:     []int{3, 14, 15, 9},
			m:        50,
			expected: 48,
		},
		{
			name:     "fhooooo",
			arry:     []int{16, 11, 2},
			expected: 21,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := darts(tc.arry, tc.m)

			if actual != tc.expected {
				t.Errorf("\nexpected is %v\nbut actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_festival(t *testing.T) {
	testCases := []struct {
		name     string
		a, b, c  []int
		expected int
	}{
		{
			name:     "fooo",
			a:        []int{1, 1, 1},
			b:        []int{2, 2, 2},
			c:        []int{3, 3, 3},
			expected: 27,
		},
		{
			name:     "fhooooo",
			a:        []int{1, 5},
			b:        []int{2, 4},
			c:        []int{3, 6},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := festival(tc.a, tc.b, tc.c)

			if actual != tc.expected {
				t.Errorf("\nexpected is %v\nbut actual is %v", tc.expected, actual)
			}
		})
	}
}

func Test_compressedMap(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		expected []int
	}{
		{
			name:     "fooo",
			arry:     []int{12, 43, 7, 15, 9},
			expected: []int{2, 4, 0, 3, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := compressedMap(tc.arry)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\nexpected is %v\nbut actual is %v", tc.expected, actual)
			}
		})
	}
}

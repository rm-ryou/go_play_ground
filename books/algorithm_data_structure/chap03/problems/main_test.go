package main

import "testing"

func Test_problem32(t *testing.T) {
	arry := []int{1, 1, 2, 3, 4, 5}
	n := 1
	expected := 2

	t.Run("arryに含まれるnの数を出力すること", func(t *testing.T) {
		actual := problem32(arry, n)

		if actual != expected {
			t.Errorf("expected is %v. but acutal is %v", expected, actual)
		}
	})
}

func Test_fetchSecondMinValue(t *testing.T) {
	arry := []int{1, 2, 3, 4, 5}
	expected := 2

	actual := fetchSecondMinValue(arry)
	if actual != expected {
		t.Errorf("expected is %v. but acutal is %v", expected, actual)
	}
}

func Test_fetchMaxDiff(t *testing.T) {
	arry := []int{1, 2, 3, 4, 5}
	expected := 4

	actual := fetchMaxDiff(arry)
	if actual != expected {
		t.Errorf("expected is %v. but acutal is %v", expected, actual)
	}
}

func Test_shiftOnly(t *testing.T) {
	testCases := []struct {
		name     string
		arry     []int
		expected int
	}{
		{
			name:     "2で一回のみ割り切れる",
			arry:     []int{2, 4, 6, 2},
			expected: 1,
		},
		{
			name:     "2で一回も割り切れない",
			arry:     []int{1, 2, 3, 5},
			expected: 0,
		},
		{
			name:     "2で2回割り切れる",
			arry:     []int{4, 8, 12},
			expected: 2,
		},
		{
			name:     "return 8",
			arry:     []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := shiftOnly(tc.arry)

			if actual != tc.expected {
				t.Errorf("expected is %v. but acutal is %v", tc.expected, actual)
			}
		})
	}
}

func Test_numOfCreateKWithTorio(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{
			name:     "return 6",
			n:        2,
			k:        2,
			expected: 6,
		},
		{
			name:     "return 1",
			n:        15,
			k:        5,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := numOfCreateKWithTorio(tc.n, tc.k)

			if actual != tc.expected {
				t.Errorf("expected is %v. but acutal is %v", tc.expected, actual)
			}
		})
	}
}

func Test_manyExpress(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "return 176",
			s:        "125",
			expected: 176,
		},
		{
			name:     "return 12656242944",
			s:        "9999999999",
			expected: 12656242944,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := manyExpress(tc.s)

			if actual != tc.expected {
				t.Errorf("expected is %v. but acutal is %v", tc.expected, actual)
			}
		})
	}
}

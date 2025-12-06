package main

import "testing"

func Test_Flog(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		ary  []int
		want int
	}{
		{
			name: "return 30",
			n:    4,
			ary:  []int{10, 30, 40, 20},
			want: 30,
		},
		{
			name: "return 0",
			n:    2,
			ary:  []int{10, 10},
			want: 0,
		},
		{
			name: "return 40",
			n:    6,
			ary:  []int{30, 10, 60, 10, 60, 50},
			want: 40,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Flog(tc.n, tc.ary)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func TestKnapsack(t *testing.T) {
	testCases := []struct {
		name string
		N    int
		W    int
		vAry []int
		wAry []int
		want int
	}{
		{
			name: "return 94",
			N:    6,
			W:    9,
			vAry: []int{3, 2, 6, 1, 3, 85},
			wAry: []int{2, 1, 3, 2, 1, 5},
			want: 94,
		},
		{
			name: "return 16",
			N:    3,
			W:    10,
			vAry: []int{15, 10, 6},
			wAry: []int{9, 6, 4},
			want: 16,
		},
	}

	for _, tc := range testCases {
		act := Knapsack(tc.N, tc.W, tc.wAry, tc.vAry)
		if act != tc.want {
			t.Errorf("want: %d, act: %d", tc.want, act)
		}
	}
}

func Test_EditDistance(t *testing.T) {
	testCases := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "return 6",
			a:    "logistic",
			b:    "algorithm",
			want: 6,
		},
		{
			name: "return 4",
			a:    "FOOD",
			b:    "MONEY",
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := EditDistance(tc.a, tc.b)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

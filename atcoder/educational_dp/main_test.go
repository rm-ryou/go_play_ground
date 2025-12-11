package main

import "testing"

func Test_Flog1(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "return 30",
			height: []int{10, 30, 40, 20},
			want:   30,
		},
		{
			name:   "return 0",
			height: []int{10, 10},
			want:   0,
		},
		{
			name:   "return 40",
			height: []int{30, 10, 60, 10, 60, 50},
			want:   40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := Flog1(len(tt.height), tt.height)
			if act != tt.want {
				t.Errorf("want: %d, act: %d", tt.want, act)
			}
		})
	}
}

func Test_Flog2(t *testing.T) {
	tests := []struct {
		name   string
		K      int
		height []int
		want   int
	}{
		{
			name:   "return 30",
			K:      3,
			height: []int{10, 30, 40, 50, 20},
			want:   30,
		},
		{
			name:   "return 20",
			K:      1,
			height: []int{10, 20, 10},
			want:   20,
		},
		{
			name:   "return 0",
			K:      100,
			height: []int{10, 10},
			want:   0,
		},
		{
			name:   "return 40",
			K:      4,
			height: []int{40, 10, 20, 70, 80, 10, 20, 70, 80, 60},
			want:   40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := Flog2(len(tt.height), tt.K, tt.height)
			if act != tt.want {
				t.Errorf("want: %d, act: %d", tt.want, act)
			}
		})
	}
}

func Test_Vacation(t *testing.T) {
	tests := []struct {
		name    string
		N       int
		a, b, c []int
		want    int
	}{
		{
			name: "return 210",
			N:    3,
			a:    []int{10, 20, 30},
			b:    []int{40, 50, 60},
			c:    []int{70, 80, 90},
			want: 210,
		},
		{
			name: "return 100",
			N:    1,
			a:    []int{100},
			b:    []int{10},
			c:    []int{1},
			want: 100,
		},
		{
			name: "return 46",
			N:    7,
			a:    []int{6, 8, 2, 7, 4, 2, 7},
			b:    []int{7, 8, 5, 8, 6, 3, 5},
			c:    []int{8, 3, 2, 6, 8, 4, 1},
			want: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := Vacation(tt.N, tt.a, tt.b, tt.c)
			if act != tt.want {
				t.Errorf("want: %d, act: %d", tt.want, act)
			}
		})
	}
}

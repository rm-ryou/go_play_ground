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

package problem

import "testing"

func Test_Vacation(t *testing.T) {
	testCases := []struct {
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Vacation(tc.N, tc.a, tc.b, tc.c)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_IsCreateValueFromAry(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		w    int
		want bool
	}{
		{
			name: "aryの要素からwを作成できる時、trueを返す",
			ary:  []int{1, 2, 3, 5, 11},
			w:    10,
			want: true,
		},
		{
			name: "aryの要素からwを作成できない時、falseを返す",
			ary:  []int{1, 5, 8, 11},
			w:    10,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsCreateValueFromAry(tc.w, tc.ary)
			if act != tc.want {
				t.Errorf("want: %t, got %t", tc.want, act)
			}
		})
	}
}

func Test_Contest(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		want int
	}{
		{
			name: "return 7",
			ary:  []int{2, 3, 5},
			want: 7,
		},
		{
			name: "return 11",
			ary:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: 11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := Context(tc.ary)
			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

func Test_IsCreateValueFromAryUnderK(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		K    int
		w    int
		want bool
	}{
		{
			name: "aryの要素からK個以下の要素を用いてwを作成できる時、trueを返す",
			ary:  []int{1, 2, 3, 5, 11},
			K:    3,
			w:    10,
			want: true,
		},
		{
			name: "aryの要素からK個以下の要素を用いてwを作成できない時、falseを返す",
			ary:  []int{1, 2, 3, 5, 11},
			K:    2,
			w:    10,
			want: false,
		},
		{
			name: "aryの要素からwを作成できない時、falseを返す",
			ary:  []int{1, 5, 8, 11},
			K:    3,
			w:    10,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsCreateValueFromAryUnderK(tc.K, tc.w, tc.ary)
			if act != tc.want {
				t.Errorf("want: %t, got %t", tc.want, act)
			}
		})
	}
}

func Test_IsCreateValueFromAryUnlimited(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		w    int
		want bool
	}{
		{
			name: "aryの要素から重複ありでwを作成できる時、trueを返す",
			ary:  []int{2, 5, 11},
			w:    8,
			want: true,
		},
		{
			name: "aryの要素から重複ありでwを作成できない時、falseを返す",
			ary:  []int{5, 8, 11},
			w:    9,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsCreateValueFromAryUnlimited(tc.w, tc.ary)
			if act != tc.want {
				t.Errorf("want: %t, got %t", tc.want, act)
			}
		})
	}
}

func Test_IsCreateValueFromAryLimited(t *testing.T) {
	testCases := []struct {
		name string
		ary  []int
		lim  []int
		w    int
		want bool
	}{
		{
			name: "aryの要素からlim以下までの重複ありでwを作成できる時、trueを返す",
			ary:  []int{2, 5, 11},
			lim:  []int{3, 2, 1},
			w:    8,
			want: true,
		},
		{
			name: "aryの要素からlim以下までの重複ありでwを作成できない時、falseを返す",
			ary:  []int{2, 5, 11},
			lim:  []int{1, 1, 1},
			w:    8,
			want: false,
		},
		{
			name: "aryの要素からlimi以下までの重複ありでwを作成できない時、falseを返す",
			ary:  []int{5, 8, 11},
			lim:  []int{1, 2, 3},
			w:    9,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := IsCreateValueFromAryLimited(tc.w, tc.ary, tc.lim)
			if act != tc.want {
				t.Errorf("want: %t, got %t", tc.want, act)
			}
		})
	}
}

func Test_LCS(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want string
		// want int
	}{
		{
			name: "return lcs",
			s:    "axyb",
			t:    "abyxb",
			want: "ayb",
			// want: 3,
		},
		{
			name: "return lcs",
			s:    "aa",
			t:    "xayaz",
			want: "aa",
			// want: 2,
		},
		{
			name: "return lcs",
			s:    "a",
			t:    "z",
			want: "",
			// want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := LCS(tt.s, tt.t)
			if act != tt.want {
				t.Errorf("want: %v, act: %v", tt.want, act)
			}
		})
	}
}

func Test_SectionDivision(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		ary  []int
		want float64
	}{
		{
			name: "return 20",
			N:    5,
			M:    3,
			ary:  []int{9, 1, 2, 3, 9},
			want: 20,
		},
		{
			name: "return 8.5",
			N:    4,
			M:    1,
			ary:  []int{14, 4, 9, 7},
			want: 8.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := SectionDivision(tt.N, tt.M, tt.ary)
			if act != tt.want {
				t.Errorf("want: %v, act: %v", tt.want, act)
			}
		})
	}
}

package unionfind

import "testing"

func TestIsSame(t *testing.T) {
	type testCase struct {
		name     string
		size     int
		initData []int
		x, y     int
		want     bool
	}

	testCases := []testCase{
		{
			name:     "要素が同じ木に属する時、trueを返す",
			size:     3,
			initData: []int{-1, -1, 1},
			x:        1,
			y:        2,
			want:     true,
		},
		{
			name:     "要素が異なる木に属する時、trueを返す",
			size:     3,
			initData: []int{-1, -1, -1},
			x:        1,
			y:        2,
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uf := NewUnionFind(tc.size)
			uf.Parent = tc.initData

			act := uf.IsSame(tc.x, tc.y)
			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

func TestUnite(t *testing.T) {
	type testCase struct {
		name     string
		size     int
		initData []int
		x, y     int
		want     bool
	}

	testCases := []testCase{
		{
			name:     "異なる木を結合すること",
			size:     3,
			initData: []int{-1, -1, -1},
			x:        1,
			y:        2,
			want:     true,
		},
		{
			name:     "元々同じ木に属する時、falseを返すこと",
			size:     3,
			initData: []int{-1, -1, 1},
			x:        1,
			y:        2,
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uf := NewUnionFind(tc.size)
			uf.Parent = tc.initData

			act := uf.Unite(tc.x, tc.y)
			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}

			if !uf.IsSame(tc.x, tc.y) {
				t.Errorf("not merged group")
			}
		})
	}
}

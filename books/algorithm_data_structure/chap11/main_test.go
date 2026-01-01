package main

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

func TestSizeOfGroup(t *testing.T) {
	N := 13
	M := 14

	// 1, 2
	// 1, 3
	// 1, 4
	// 2, 5
	// 2, 6
	// 3, 6
	// 4, 7
	// 5, 6
	// 5, 8
	// 6, 7
	// 7, 9
	// 8, 9
	// 11, 13
	// 12, 13
	a := []int{1, 1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 11, 12}
	b := []int{2, 3, 4, 5, 6, 6, 7, 6, 8, 7, 9, 9, 13, 13}

	t.Run("グラフの連結成分の個数を取得する", func(t *testing.T) {
		want := 3
		uf := NewUnionFind(N)
		act := uf.SizeOfGroup(M, a, b)

		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

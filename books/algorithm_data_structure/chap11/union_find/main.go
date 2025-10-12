package main

import "fmt"

type UnionFind struct {
	par, siz []int
}

func NewUF(n int) *UnionFind {
	par, siz := make([]int, n), make([]int, n)
	return &UnionFind{
		par: par,
		siz: siz,
	}
}

func (uf *UnionFind) root(x int) int {
	if uf.par[x] == -1 {
		return x
	} else {
		uf.par[x] = uf.root(uf.par[x])
		return uf.par[x]
	}
}

func (uf *UnionFind) isSame(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *UnionFind) unite(x, y int) bool {
	x = uf.root(x)
	y = uf.root(y)

	if x == y {
		return false
	}

	if uf.siz[x] < uf.siz[y] {
		x, y = y, x
	}

	uf.par[y] = x
	uf.siz[x] += uf.siz[y]
	return true
}

func (uf *UnionFind) size(x int) int {
	return uf.siz[uf.root(x)]
}

func main() {
	uf := NewUF(7)

	uf.unite(1, 2)
	uf.unite(2, 3)
	uf.unite(5, 6)
	fmt.Println(uf.isSame(1, 3))
	fmt.Println(uf.isSame(2, 5))

	uf.unite(1, 6)
	fmt.Println(uf.isSame(2, 5))
}

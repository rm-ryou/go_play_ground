package main

type UnionFind struct {
	Parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	return &UnionFind{
		Parent: parent,
		size:   make([]int, n),
	}
}

func (uf *UnionFind) root(x int) int {
	if uf.Parent[x] == -1 {
		return x
	}

	uf.Parent[x] = uf.root(uf.Parent[x])
	return uf.Parent[x]
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *UnionFind) Unite(x, y int) bool {
	rootX, rootY := uf.root(x), uf.root(y)
	if rootX == rootY {
		return false
	}

	if uf.Size(rootX) < uf.Size(rootY) {
		rootX, rootY = rootY, rootX
	}

	uf.Parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]

	return true
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.root(x)]
}

func (uf *UnionFind) SizeOfGroup(M int, a, b []int) int {
	for i := range M {
		uf.Unite(a[i]-1, b[i]-1)
	}

	var res int
	for i := range len(uf.Parent) {
		if uf.root(i) == i {
			res++
		}
	}
	return res
}

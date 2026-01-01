package unionfind

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

func (uf *UnionFind) Root(x int) int {
	if uf.Parent[x] == -1 {
		return x
	}

	uf.Parent[x] = uf.Root(uf.Parent[x])
	return uf.Parent[x]
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Unite(x, y int) bool {
	rootX, rootY := uf.Root(x), uf.Root(y)
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
	return uf.Parent[uf.Root(x)]
}

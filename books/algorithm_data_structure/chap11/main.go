package main

import unionfind "example.com/pkg/union_find"

func SizeOfGroup(N, M int, a, b []int) int {
	uf := unionfind.NewUnionFind(N)

	for i := range M {
		uf.Unite(a[i]-1, b[i]-1)
	}

	var res int
	for i := range len(uf.Parent) {
		if uf.Root(i) == i {
			res++
		}
	}
	return res
}

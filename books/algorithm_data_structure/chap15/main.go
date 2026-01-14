package main

import (
	"slices"

	"example.com/pkg/unionfind"
)

type Edge struct {
	a, b   int
	weight int
}

func Kruskal(N, M int, a, b, w []int) int {
	edges := make([]Edge, M)
	for i := range edges {
		e := Edge{
			a:      a[i],
			b:      b[i],
			weight: w[i],
		}
		edges[i] = e
	}

	slices.SortFunc(edges, func(a, b Edge) int {
		if a.weight > b.weight {
			return 1
		} else {
			return -1
		}
	})

	var res int
	uf := unionfind.NewUnionFind(N)
	for i := range M {
		if uf.IsSame(edges[i].a, edges[i].b) {
			continue
		}

		res += edges[i].weight
		uf.Unite(edges[i].a, edges[i].b)
	}

	return res
}

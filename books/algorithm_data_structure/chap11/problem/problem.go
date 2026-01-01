package problem

import (
	unionfind "example.com/pkg/union_find"
)

func Bridge(N, M int, A, B []int) int {
	var res int

	for i := range M {
		uf := unionfind.NewUnionFind(N)
		for j := range M {
			if i == j {
				continue
			}

			uf.Unite(A[j]-1, B[j]-1)
		}

		var unit int
		for u := range N {
			if uf.Root(u) == u {
				unit++
			}
		}

		if unit > 1 {
			res++
		}
	}

	return res
}

func Decayed(N, M int, A, B []int) []int {
	res := make([]int, M)

	curInconvenience := (N * (N - 1)) / 2
	for i := range M {
		res[M-(1+i)] = curInconvenience
		uf := unionfind.NewUnionFind(N)

		a := A[M-(1+i)] - 1
		b := B[M-(1+i)] - 1

		if uf.IsSame(a, b) {
			continue
		}

		curInconvenience -= uf.Size(a) * uf.Size(b)
		uf.Unite(a, b)

	}
	return res
}

package problem

import (
	"slices"
	"sort"
)

func Pair(a, b []int) int {
	var count int

	sort.Ints(a)
	sort.Ints(b)

	for _, v := range b {
		if a[count] < v {
			count++
		}
	}

	return count
}

type Cod struct {
	X, Y int
}

func sortCodByX(c []Cod) []Cod {
	slices.SortFunc(c, func(a, b Cod) int {
		if a.X > b.X {
			return 1
		} else {
			return -1
		}
	})

	return c
}

func PlainPoints(n int, a, b []Cod) int {
	var count int

	sortedB := sortCodByX(b)
	isUsed := make([]bool, n)

	for _, vb := range sortedB {

		hitId := -1
		maxYOfA := -1
		for j, v := range a {
			if isUsed[j] {
				continue
			}

			if v.X >= vb.X {
				continue
			}
			if v.Y < maxYOfA {
				continue
			}

			maxYOfA = v.Y
			hitId = j
		}

		if hitId != -1 {
			count++
			isUsed[hitId] = true
		}
	}

	return count
}

func Megalo(n int, a, b []int) bool {
	flg := true

	dAry := make([][]int, n)
	for i := range dAry {
		dAry[i] = make([]int, 2)
		dAry[i][0] = a[i] // 終えるのに要する時間
		dAry[i][1] = b[i] // 締め切り
	}

	slices.SortFunc(dAry, func(a, b []int) int {
		if a[1] > b[1] {
			return 1
		} else {
			return -1
		}
	})

	var finishTime int
	for _, v := range dAry {
		finishTime += v[0]
		if finishTime > v[1] {
			flg = false
		}
	}

	return flg
}

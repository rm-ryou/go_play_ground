package problem

import (
	"math"
	"sort"
)

func lowerBound(ary []int, key int) int {
	idx := sort.Search(len(ary), func(i int) bool {
		return ary[i] >= key
	})
	return idx
}

func upperBound(ary []int, key int) int {
	idx := sort.Search(len(ary), func(i int) bool {
		return ary[i] > key
	})
	return idx
}

func Compression(a []int) []int {
	size := len(a)

	c := make([]int, size)
	copy(c, a)
	sort.Ints(c)

	res := make([]int, size)
	for i := range size {
		res[i] = lowerBound(c, a[i])
	}

	return res
}

func Snuke(a, b, c []int) int {
	var res int
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	for i := range b {
		an := lowerBound(a, b[i])
		cn := len(c) - upperBound(c, b[i])
		res += an * cn
	}

	return res
}

func Dart(m int, ary []int) int {
	var res int
	ary = append(ary, 0)
	var s []int
	for i := range ary {
		for j := range ary {
			s = append(s, ary[i]+ary[j])
		}
	}

	sort.Ints(s)
	for _, v := range s {
		idx := upperBound(s, m-v)
		if idx == 0 {
			continue
		}
		res = int(math.Max(float64(res), float64(v+s[idx-1])))
	}

	return res
}

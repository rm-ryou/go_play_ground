package main

import (
	"math"
	"sort"
)

func cows(arry []int, m int) int {

	left, right := 0, math.MaxInt
	for right-left > 1 {
		mid := (left + right) / 2

		count := 1
		prev := 0
		for i := range len(arry) {
			if arry[i]-arry[prev] >= mid {
				count++
				prev = i
			}
		}

		if count >= m {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func darts(arry []int, m int) int {
	var tmp []int
	for _, v := range arry {
		for _, v2 := range arry {
			tmp = append(tmp, v+v2)
		}
	}

	binarySearch := func(key int) int {
		left, right := 0, len(tmp)

		for right-left > 1 {
			mid := left + (right-left)/2
			if tmp[mid] == key {
				return mid
			} else if tmp[mid] > key {
				right = mid
			} else {
				left = mid
			}
		}
		return left
	}

	res := 0
	sort.Ints(tmp)
	for _, v := range tmp {
		maxIdx := binarySearch(m - v)
		tmpValue := v + tmp[maxIdx-1]
		if res < tmpValue {
			res = tmpValue
		}
	}
	return res
}

// ソート済みのスライスarryのうち、key以下要素の最大の添字
func upperMaxIdxOfKey(arry []int, key int) int {
	left, right := 0, len(arry)

	for right-left > 1 {
		mid := left + (right-left)/2
		if arry[mid] > key {
			right = mid
		} else {
			left = mid
		}
	}

	if arry[right-1] > key {
		return right - 1
	} else {
		return right
	}
}

// ソート済みのスライスarryのうち、keyより小さい要素の最大の添字
func lowerMaxIdxOfKey(arry []int, key int) int {
	left, right := 0, len(arry)

	for right-left > 1 {
		mid := left + (right-left)/2
		if arry[mid] >= key {
			right = mid
		} else {
			left = mid
		}
	}
	return left + 1
}

func festival(a, b, c []int) int {
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	res := 0
	for _, v := range b {
		countOfA := lowerMaxIdxOfKey(a, v)
		countOfC := len(c) - upperMaxIdxOfKey(c, v)
		res = res + (countOfA * countOfC)
	}

	return res
}

func binarySearchWithPair(arry []int, key int) int {
	left, right := 0, len(arry)

	res := -1
	for right-left >= 1 {
		mid := left + (right-left)/2
		if arry[mid] == key {
			res = mid
			break
		} else if arry[mid] > key {
			right = mid
		} else {
			left = mid
		}
	}
	return res
}

func compressedMap(arry []int) []int {
	dist := make([]int, len(arry))
	copy(dist, arry)

	res := make([]int, len(arry))
	sort.Ints(dist)

	for i := range len(arry) {
		res[i] = binarySearchWithPair(dist, arry[i])
	}
	return res
}

func main() {}

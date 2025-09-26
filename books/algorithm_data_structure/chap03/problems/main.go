package main

import (
	"math"
)

func problem32(arry []int, n int) int {
	res := 0
	for _, v := range arry {
		if v == n {
			res++
		}
	}
	return res
}

func fetchSecondMinValue(arry []int) int {
	minValue, secondMinValue := math.MaxInt, math.MaxInt

	for _, v := range arry {
		if v < minValue {
			secondMinValue = minValue
			minValue = v
		} else if v < secondMinValue {
			secondMinValue = v
		}
	}

	return secondMinValue
}

func fetchMaxDiff(arry []int) int {
	maxDiff := 0

	for i, v := range arry {
		for _, v2 := range arry[i+1:] {
			diff := int(math.Abs(float64(v - v2)))
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	return maxDiff
}

func shiftOnly(arry []int) int {
	res := math.MaxInt

	for _, bit := range arry {
		// 最後の1bitが立つまでシフトを行い、その中で最小の値を取得しておく
		canDivideWith2 := 0
		for bit&0x1 == 0 {
			canDivideWith2++
			bit = bit >> 1
		}
		if canDivideWith2 < res {
			res = canDivideWith2
		}
	}
	return res
}

func numOfCreateKWithTorio(n, k int) int {
	res := 0

	for v1 := range k + 1 {
		for v2 := range k + 1 {
			v3 := n - (v1 + v2)
			if v3 >= 0 && v3 <= k {
				res++
			}
		}
	}
	return res
}

func slice2num(arry []int) int {
	if len(arry) == 1 {
		return arry[0]
	}

	res := 0

	for _, v := range arry {
		res *= 10
		res += v
	}
	return res
}

func manyExpress(s string) int {
	res := 0
	// +の入れ方は2^(len(s) - 1)
	// sliceに戻して、各要素を2進数でフラグが立っている個所で分ける
	// []int {1, 2, 5}
	// 00 = 125
	// 01 = 1 + 25
	// 10 = 12 + 5
	// 11 = 1 + 2 + 5

	arry := make([]int, len(s))
	for i, c := range s {
		if c >= '0' && c <= '9' {
			arry[i] = int(c - '0')
		}
	}

	for bit := range 1 << (len(s) - 1) {
		if bit == 0 {
			res += slice2num(arry)
			continue
		}

		i, j := 0, 0
		for range bit {
			j++
			// 2. (i, j]のスライスをたす
			// 3. iの位置をjにする
			if bit&0x1 == 1 {
				res += slice2num(arry[i:j])
				i = j
			}
			bit = bit >> 1
		}

		res += slice2num(arry[i:])
	}
	return res
}

func main() {}

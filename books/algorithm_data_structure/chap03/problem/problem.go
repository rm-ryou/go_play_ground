package problem

import (
	"math"
)

func CountValue(a []int, key int) int {
	count := 0

	for _, v := range a {
		if v == key {
			count++
		}
	}

	return count
}

func FetchSecondMinValue(a []int) int {
	firstMinValue := math.MaxInt
	secondMinValue := math.MaxInt

	for _, v := range a {
		if v < firstMinValue {
			secondMinValue = firstMinValue
			firstMinValue = v
		} else if v < secondMinValue {
			secondMinValue = v
		}
	}

	return secondMinValue
}

func CalcMaxDiff(a []int) int {
	maxValue := math.MinInt
	minValue := math.MaxInt

	for _, v := range a {
		if v < minValue {
			minValue = v
		} else if v > maxValue {
			maxValue = v
		}
	}

	return maxValue - minValue
}

func CalcMaxDivide(a []int) int {
	res := math.MaxInt

	for _, v := range a {
		count := 0
		for i := range v {
			if v&(1<<i) != 0 {
				break
			}
			count++
		}
		if res > count {
			res = count
		}
	}

	return res
}

func SumOfThreeNum(n, k int) int {
	count := 0

	for i := range k + 1 {
		for j := range k + 1 {
			diff := n - (i + j)
			if diff >= 0 && diff <= k {
				count++
			}
		}
	}

	return count
}

func ManyExpre(s string) int {
	res := 0
	ary := make([]int, len(s))

	for i, c := range s {
		ary[i] = int(c - '0')
	}

	n := len(ary)
	for bit := range 1 << (n - 1) {
		sum := 0
		for i := range n - 1 {
			sum *= 10
			sum += ary[i]

			if bit&(1<<i) != 0 {
				res += sum
				sum = 0
			}
		}

		sum *= 10
		sum += ary[n-1]
		res += sum
	}

	return res
}

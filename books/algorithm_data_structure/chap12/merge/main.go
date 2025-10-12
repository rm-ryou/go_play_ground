package main

import "fmt"

func mergeSort(a []int, left, right int) {
	if right-left == 1 {
		return
	}

	mid := left + (right-left)/2

	mergeSort(a, left, mid)
	mergeSort(a, mid, right)

	var buf []int
	for i := left; i < mid; i++ {
		buf = append(buf, a[i])
	}
	for i := right - 1; i >= mid; i-- {
		buf = append(buf, a[i])
	}

	leftIdx := 0
	rightIdx := len(buf) - 1
	for i := left; i < right; i++ {
		if buf[leftIdx] <= buf[rightIdx] {
			a[i] = buf[leftIdx]
			leftIdx += 1
		} else {
			a[i] = buf[rightIdx]
			rightIdx--
		}
	}
}

func main() {
	arry := []int{8, 7, 6, 4, 1, 2, 5, 9, 13}
	mergeSort(arry, 0, len(arry)-1)
	fmt.Println(arry)
}

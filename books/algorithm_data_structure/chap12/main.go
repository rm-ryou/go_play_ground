package main

func InsertionSort(data []int) []int {
	clone := make([]int, len(data))
	copy(clone, data)

	for i := range clone {
		for j := i + 1; j < len(clone); j++ {
			if clone[i] > clone[j] {
				clone[i], clone[j] = clone[j], clone[i]
			}
		}
	}

	return clone
}

func merge(data []int, left, mid, right int) {
	var buf []int
	for i := left; i < mid; i++ {
		buf = append(buf, data[i])
	}
	for i := right - 1; i >= mid; i-- {
		buf = append(buf, data[i])
	}

	leftIdx, rightIdx := 0, len(buf)-1
	for i := left; i < right; i++ {
		if buf[leftIdx] <= buf[rightIdx] {
			data[i] = buf[leftIdx]
			leftIdx++
		} else {
			data[i] = buf[rightIdx]
			rightIdx--
		}
	}
}

func doMergeSort(data []int, left, right int) {
	if right-left == 1 {
		return
	}

	mid := left + (right-left)/2

	doMergeSort(data, left, mid)
	doMergeSort(data, mid, right)

	merge(data, left, mid, right)
}

func MergeSort(data []int) []int {
	clone := make([]int, len(data))
	copy(clone, data)

	doMergeSort(clone, 0, len(clone))

	return clone
}

func doQuickSort(data []int, left, right int) {
	if right-left <= 1 {
		return
	}

	midIdx := (left + right) / 2
	mid := data[midIdx]
	data[midIdx], data[right-1] = data[right-1], data[midIdx]

	base := left
	for i := left; i < right-1; i++ {
		if data[i] < mid {
			data[base], data[i] = data[i], data[base]
			base++
		}
	}
	data[base], data[right-1] = data[right-1], data[base]

	doQuickSort(data, left, base)
	doQuickSort(data, base+1, right)
}

func QuickSort(data []int) []int {
	clone := make([]int, len(data))
	copy(clone, data)

	doQuickSort(clone, 0, len(clone))

	return clone
}

// iを根とする部分木
func Heapify(data []int, i, n int) {
	child := i*2 + 1
	if child >= n {
		return
	}

	if child+1 < n && data[child+1] > data[child] {
		child++
	}

	if data[child] <= data[i] {
		return
	}

	data[i], data[child] = data[child], data[i]

	Heapify(data, child, n)
}

func HeapSort(data []int) []int {
	clone := make([]int, len(data))
	copy(clone, data)

	for i := len(clone)/2 - 1; i >= 0; i-- {
		Heapify(clone, i, len(clone))
	}
	for i := len(clone) - 1; i > 0; i-- {
		clone[0], clone[i] = clone[i], clone[0]
		Heapify(clone, 0, i)
	}

	return clone
}

func BucketSort(data []int) []int {
	MaxElemNumber := 1000
	num := make([]int, MaxElemNumber)
	for i := range len(data) {
		num[data[i]]++
	}

	sum := make([]int, MaxElemNumber)
	for i := 1; i < MaxElemNumber; i++ {
		sum[i] = sum[i-1] + num[i]
	}

	res := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		idx := sum[data[i]] - 1
		res[idx] = data[i]
	}

	return res
}

package main

import "fmt"

func heapify(a []int, i, n int) {
	child1 := i*2 + 1
	if child1 >= n {
		return
	}

	if child1+1 < n && a[child1+1] > a[child1] {
		child1++
	}

	if a[child1] <= a[i] {
		return
	}

	a[i], a[child1] = a[child1], a[i]

	heapify(a, child1, n)
}

func heapSort(a []int) {
	n := len(a)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, i, n)
	}
	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, 0, i)
	}
}

func main() {
	arry := []int{8, 7, 6, 4, 1, 2, 5, 9, 13, 12}
	heapSort(arry)
	fmt.Println(arry)
}

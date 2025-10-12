package main

import "fmt"

func quick(a []int, left, right int) {
	if (right - left) <= 1 {
		return
	}

	pivoIdx := (left + right) / 2
	pivo := a[pivoIdx]

	a[pivoIdx], a[right-1] = a[right-1], a[pivoIdx]

	i := left
	for j := left; j < right-1; j++ {
		if a[j] < pivo {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[right-1] = a[right-1], a[i]

	quick(a, left, 1)
	quick(a, i+1, right)
}

func main() {
	arry := []int{8, 7, 6, 4, 1, 2, 5, 9, 13, 12}
	quick(arry, 0, len(arry))
	fmt.Println(arry)
}

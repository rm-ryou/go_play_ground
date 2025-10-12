package main

import "fmt"

func Insertion(arry []int) {
	n := len(arry)

	for i := range n {
		v := arry[i]
		j := i
		for ; j > 0; j-- {
			if arry[j-1] > v {
				arry[j] = arry[j-1]
			} else {
				break
			}
		}
		arry[j] = v
	}
}

func main() {
	arry := []int{8, 7, 6, 4, 1, 2, 5, 9, 13, 12}
	Insertion(arry)
	fmt.Println(arry)
}

package main

import "fmt"

type heap []int

func (h *heap) push(x int) {
	*h = append(*h, x)
	i := len(*h) - 1
	for i > 0 {
		parentIdx := (i - 1) / 2
		if (*h)[parentIdx] >= x {
			break
		}
		(*h)[i] = (*h)[parentIdx]
		i = parentIdx
	}
	(*h)[i] = x
}

func (h *heap) pop() {
	if len(*h) == 0 {
		return
	}

	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-2]
	i := 0
	for i*2+1 < len(*h) {
		child1, child2 := i*2+1, i*2+2
		if child2 < len(*h) && (*h)[child2] > (*h)[child1] {
			child1 = child2
		}
		if (*h)[child1] <= x {
			break
		}
		(*h)[i] = (*h)[child1]
		i = child1
	}
	(*h)[i] = x
}

func (h *heap) top() int {
	if len(*h) != 0 {
		return (*h)[0]
	}
	return -1
}

func main() {
	var h heap
	h.push(5)
	h.push(3)
	h.push(7)
	h.push(1)

	fmt.Println(h.top())
	h.pop()
	fmt.Println(h.top())
	h.push(11)
	fmt.Println(h.top())
}

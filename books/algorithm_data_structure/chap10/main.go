package main

import (
	"errors"
	"fmt"
)

var ErrDataIsEmpty = errors.New("data is empty")

type Heap struct {
	data []int
}

func (h *Heap) Push(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1
	for i > 0 {
		parentIdx := (i - 1) / 2
		if h.data[parentIdx] >= val {
			break
		}

		h.data[i] = h.data[parentIdx]
		i = parentIdx
		h.data[i] = val
	}
}

func (h *Heap) Top() int {
	if len(h.data) == 0 {
		return -1
	}
	return h.data[0]
}

func (h *Heap) Pop() error {
	if len(h.data) == 0 {
		return ErrDataIsEmpty
	}

	tmp := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	i := 0
	for i*2+1 < len(h.data) {
		child1Idx, child2Idx := i*2+1, i*2+2
		if child2Idx < len(h.data) && h.data[child2Idx] > h.data[child1Idx] {
			child1Idx = child2Idx
		}

		if h.data[child1Idx] <= tmp {
			break
		}
		h.data[i] = h.data[child1Idx]
		i = child1Idx
	}
	h.data[i] = tmp

	return nil
}

func main() {
	var h Heap

	h.Push(5)
	h.Push(3)
	h.Push(7)
	h.Push(1)

	fmt.Println(h.data)

	fmt.Println(h.Top())
	h.Pop()
	fmt.Println(h.Top())
	h.Push(11)
	fmt.Println(h.Top())
}

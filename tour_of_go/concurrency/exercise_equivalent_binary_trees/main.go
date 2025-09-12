package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left *Tree
// 	value int
// 	Right *Tree
// }

func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkRecursive(t.Left, ch)
	ch <- t.Value
	walkRecursive(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			break
		}

		if ok1 && v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

func Walk(t *tree.Tree, ch chan int) {
	walkRecv(t, ch)
	close(ch)
}

func walkRecv(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkRecv(t.Left, ch)
	ch <- t.Value
	walkRecv(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		if !ok1 {
			break
		}
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for range 10 {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}

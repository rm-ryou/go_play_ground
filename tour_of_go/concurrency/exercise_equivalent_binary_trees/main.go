package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// tree pakcage
// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walkRecv(t, ch)
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
	ch1, ch2 := make(chan int, 10), make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		if !ok1 {
			return true
		}
		v2, ok2 := <-ch2
		if ok1 && (ok2 != ok1 || v1 != v2) {
			return false
		}
	}
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

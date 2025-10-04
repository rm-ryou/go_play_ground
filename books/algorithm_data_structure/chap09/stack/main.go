package main

import "fmt"

type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic("stack is empty")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res
}

func main() {
	var s stack

	s.push(4)
	s.push(5)
	s.push(7)

	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}

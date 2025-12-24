package main

type Node struct {
	Next *Node
	Name string
}

func NewNode(name string) *Node {
	return &Node{
		Next: nil,
		Name: name,
	}
}

func (n *Node) Insert(val *Node) {
	val.Next = n.Next
	n.Next = val
}

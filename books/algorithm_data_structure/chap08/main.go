package main

type Node struct {
	Next *Node
	Prev *Node
	Name string
}

func NewNode(name string) *Node {
	return &Node{
		Next: nil,
		Prev: nil,
		Name: name,
	}
}

func (n *Node) Insert(val *Node) {
	val.Next = n.Next
	if n.Next != nil {
		n.Next.Prev = val
	}
	n.Next = val
	val.Prev = n
}

func (n *Node) Delete() {
	if n == nil {
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

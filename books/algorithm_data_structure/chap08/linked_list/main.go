package main

import "fmt"

type Node struct {
	next *Node
	name string
}

func initNode() *Node {
	return &Node{
		next: nil,
		name: "",
	}
}

func (n *Node) insert(newNode *Node) {
	newNode.next = n.next
	n.next = newNode
}

func (n *Node) printList() {
	cur := n.next
	for ; cur != nil; cur = cur.next {
		fmt.Printf("%s -> ", cur.name)
	}
	fmt.Println("<nil>")
}

func main() {
	first := initNode()

	names := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}

	for _, name := range names {
		newNode := &Node{next: nil, name: name}

		first.insert(newNode)
		first.printList()
	}
}

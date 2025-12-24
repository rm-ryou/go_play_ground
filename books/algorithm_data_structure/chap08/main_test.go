package main

import "testing"

func Test_Node(t *testing.T) {
	t.Run("終端ノードの次ノードはnilであること", func(t *testing.T) {
		node := NewNode("test")

		if node.Next != nil {
			t.Errorf("want: %v, act: %v", nil, node.Next)
		}
	})
}

func Test_Insert(t *testing.T) {
	t.Run("prevの直後にvalが挿入されること", func(t *testing.T) {
		node := NewNode("first")
		nextNode := NewNode("second")

		node.Insert(nextNode)

		if node.Next == nil {
			t.Errorf("want: %v, act: %v", nextNode, node.Next)
		}

		if node.Next.Name != nextNode.Name {
			t.Errorf("want: %s, act: %s", nextNode.Name, node.Next.Name)
		}
	})
}

func Test_Delete(t *testing.T) {
	firstNode := NewNode("first")
	secondNode := NewNode("second")
	thirdNode := NewNode("third")

	firstNode.Insert(secondNode)
	secondNode.Insert(thirdNode)

	if firstNode.Next.Next.Name != thirdNode.Name {
		t.Fatalf("failed to setup. want: %s, act: %s", thirdNode.Name, firstNode.Next.Next.Name)
	}

	secondNode.Delete()

	if firstNode.Next.Name != thirdNode.Name {
		t.Errorf("want: %s, act: %s", thirdNode.Name, firstNode.Next.Name)
	}
}

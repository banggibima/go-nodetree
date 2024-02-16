package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return & Node{Value: value, Left: nil, Right: nil}
	}

	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else {
		root.Right = insert(root.Right, value)
	}

	return root
}

func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Printf("%d ", root.Value)
		inorderTraversal(root.Right)
	}
}

func main() {
	var root *Node
	var n, value int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&value)
		root = insert(root, value)
	}

	inorderTraversal(root)
	fmt.Printf("\n")
}

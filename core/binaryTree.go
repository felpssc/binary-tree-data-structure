package core

import "fmt"

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(data int) *BinaryTree {
	if b.root != nil {
		b.root.Insert(data)

	} else {
		b.root = NewTreeNode(data)
	}
	return b
}

func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}

	PrintTree(node.Right, level+1)

	for i := 0; i < level; i++ {
		fmt.Print("   ")
	}

	fmt.Println(node.Data)

	PrintTree(node.Left, level+1)
}

func (b BinaryTree) GetRoot() *Node {
	return b.root
}

func (b *BinaryTree) CountNodes(node *Node) int {
	if node == nil {
		return 0
	}

	return b.CountNodes(node.Left) + b.CountNodes(node.Right) + 1
}

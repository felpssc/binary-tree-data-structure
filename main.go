package main

import (
	"fmt"

	"github.com/felpssc/binary-tree-data-structure/core"
)

func main() {
	tree := core.NewBinaryTree()

	var input int = 1

	for input != 0 {
		fmt.Println("Enter a number to insert in the Binary Tree (0 to stop and print result): ")
		fmt.Scanf("%d", &input)

		if input != 0 {
			tree.Insert(input)
		}
	}

	core.PrintTree(tree.GetRoot(), 1)
	fmt.Println("Nodes count: ", tree.CountNodes(tree.GetRoot()))
}

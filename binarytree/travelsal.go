package binarytree

import "fmt"

func beforeRecursion(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.val)
	beforeRecursion(node.left)
	beforeRecursion(node.right)
}

func beforeNoRecursion(node *Node) {
	//s :=
}

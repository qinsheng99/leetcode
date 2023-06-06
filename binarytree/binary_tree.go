package binarytree

type Node struct {
	val   int
	left  *Node
	right *Node
}

func newNode(v int) *Node {
	return &Node{val: v}
}

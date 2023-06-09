package binarytree

import "testing"

func TestBefore(t *testing.T) {
	n := newNode(1)
	n.left = newNode(2)
	n.right = newNode(3)
	n.left.left = newNode(4)
	n.left.right = newNode(5)

	n.right.left = newNode(6)
	n.right.right = newNode(7)

	beforeRecursion(n)
}

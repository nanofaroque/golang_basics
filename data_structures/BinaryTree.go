package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	root := &Node{val: 6, left: nil, right: nil,}

	root.left = &Node{val: 5, left: nil, right: nil}
	root.left = &Node{val: 7, left: nil, right: nil}
	//fmt.Print(root.val)

	printHelper(root)
}

func printHelper(root *Node) {
	fmt.Print(root.val)
	if root == nil {
		return
	}
	printHelper(root.left)
	printHelper(root.right)

}

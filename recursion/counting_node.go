package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func main() {
	root := &Node{1, nil}
	root.next = &Node{2, nil}
	root.next.next = &Node{3, nil}
	res := print_node(root)
	fmt.Print(res)
}

func print_node(node *Node) int {
	if node == nil {
		return 0
	}
	fmt.Print(node.val)
	return node.val + print_node(node.next)
}

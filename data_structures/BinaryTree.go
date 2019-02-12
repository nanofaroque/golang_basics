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
	root.right = &Node{val: 7, left: nil, right: nil}
	//fmt.Print(root.val)

	printHelper(root)
}

func printHelper(root *Node)  {
	fmt.Println(root.val)
	fmt.Println(root.left.val)
	fmt.Println(root.right.val)

}

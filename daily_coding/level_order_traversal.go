package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.right.left = &Node{4, nil, nil}
	m := make(map[int][]int)
	traverse(root, m, 0)
	fmt.Print(m)

}

func traverse(root *Node, m map[int][]int, level int) {
	if root==nil {return}
	if root != nil {
		m[level] = append(m[level], root.val)
	}
	traverse(root.left,m,level+1)
	traverse(root.right,m,level+1)
}

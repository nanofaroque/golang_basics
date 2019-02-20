package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*** To pass the test cases, we need some clean up code
   whenever isCousins return, make sure you put
   result=nil
 */
var result [][]int

func isCousins(root *TreeNode, x int, y int) bool {

	helper(root, x, y, 0)
	fmt.Print(result)
	if len(result) != 2 {
		return false
	}

	if result[0][1] != result[1][1] {
		return false
	}

	if result[0][0] == result[1][0] {
		return false
	}

	return true
}

func helper(root *TreeNode, x int, y int, level int) {
	// result size is 2 means, we already found the targets
	if len(result) == 2 {
		return
	}

	if root == nil {
		return
	}

	if root.Left != nil && (root.Left.Val == x || root.Left.Val == y) {
		result = append(result, []int{root.Val, level + 1})
	}

	if root.Right != nil && (root.Right.Val == x || root.Right.Val == y) {
		result = append(result, []int{root.Val, level + 1})
	}
	helper(root.Left, x, y, level+1)
	helper(root.Right, x, y, level+1)
}

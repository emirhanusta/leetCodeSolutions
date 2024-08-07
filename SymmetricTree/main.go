package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeNode := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	fmt.Println(isSymmetric(treeNode))
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return checkIsSymmetric(root.Left, root.Right)
}

func checkIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return checkIsSymmetric(left.Left, right.Right) && checkIsSymmetric(left.Right, right.Left)
}

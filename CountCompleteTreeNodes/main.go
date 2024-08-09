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
	fmt.Println(countNodes(treeNode))
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

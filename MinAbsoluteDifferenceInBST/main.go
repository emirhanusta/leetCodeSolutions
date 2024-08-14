package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeNode := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	fmt.Println(getMinimumDifference(treeNode))
}

var (
	X       int
	MinDiff int
)

func getMinimumDifference(root *TreeNode) int {
	X = math.MaxInt64
	MinDiff = math.MaxInt64
	inorder(root)
	return MinDiff
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)

	if root.Val > X {
		MinDiff = min(MinDiff, root.Val-X)
	} else {
		MinDiff = min(MinDiff, X-root.Val)
	}
	X = root.Val

	inorder(root.Right)
	return
}

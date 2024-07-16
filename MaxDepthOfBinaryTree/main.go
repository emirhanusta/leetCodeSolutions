package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	result := maxDepth(root)
	println(result)
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return 1 + maxValue(maxDepth(root.Left), maxDepth(root.Right))
	}

	return 0
}

func maxValue(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

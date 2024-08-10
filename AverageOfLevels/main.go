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
	fmt.Println(averageOfLevels(treeNode))
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var averages []float64
	currLevel := []*TreeNode{root}

	for len(currLevel) != 0 {
		var nextLevel []*TreeNode
		sum := 0.0

		for _, node := range currLevel {
			if node != nil {
				sum += float64(node.Val)
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		avg := sum / float64(len(currLevel))
		averages = append(averages, avg)
		currLevel = nextLevel
	}
	return averages
}

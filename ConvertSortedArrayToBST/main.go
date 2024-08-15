package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}).Val)
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	node := len(nums) / 2
	return &TreeNode{
		Val:   nums[node],
		Left:  sortedArrayToBST(nums[:node]),
		Right: sortedArrayToBST(nums[node+1:]),
	}
}

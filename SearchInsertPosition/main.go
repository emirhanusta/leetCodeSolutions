package main

import "fmt"

func main() {
	result := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

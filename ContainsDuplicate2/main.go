package main

import (
	"fmt"
)

func main() {
	result := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	fmt.Println(result)
}

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)

	for i, num := range nums {
		value, ok := m[num]
		if ok && i-value <= k {
			return true
		}
		m[num] = i
	}
	return false
}

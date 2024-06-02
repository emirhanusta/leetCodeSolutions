package main

import (
	"fmt"
)

func main() {
	result := merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	fmt.Println(result)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	p1, p2 := m-1, n-1

	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}

	return nums1

}

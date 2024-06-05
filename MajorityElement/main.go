package main

import "fmt"

func main() {
	result := majorityElement([]int{1, 2, 1, 3, 1, 4, 1, 5})
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	/*l := len(nums)
	i := 0
	for _, num := range nums {
		for j := 0; j < l; j++ {
			if num == nums[j] {
				i++
				if i > l/2 {
					return num
				}
			}
		}
		i = 0
	}
	return 0*/
	/*numMap := make(map[int]int)

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			numMap[num] += 1
		} else {
			numMap[num] = 1
		}
		if numMap[num] > len(nums)/2 {
			return num
		}
	}
	return 0*/

	// Boyer-Moore Voting Algorithm
	majority := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
		}
		if nums[i] == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}

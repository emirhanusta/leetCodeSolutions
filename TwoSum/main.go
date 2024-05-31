package main

import "fmt"

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		value := target - num
		if index, ok := numsMap[value]; ok {
			return []int{index, i}
		}
		numsMap[num] = i
	}
	return []int{}
}

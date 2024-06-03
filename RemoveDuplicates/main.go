package main

import "fmt"

func main() {
	result := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	duplicates := make(map[int]int)

	i := 0
	for _, num := range nums {
		if _, ok := duplicates[num]; !ok {
			duplicates[num] = i
			nums[i] = num
			i++
		}
	}
	fmt.Println(nums)
	return i
}

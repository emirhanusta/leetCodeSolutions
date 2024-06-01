package main

import "fmt"

func main() {
	result := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	k := 0
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}

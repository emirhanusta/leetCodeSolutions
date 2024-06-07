package main

import (
	"fmt"
)

func main() {
	result := lengthOfLastWord("a ")
	fmt.Println(result)
}

func lengthOfLastWord(s string) int {
	count := 0
	if len(s) == 1 {
		return len(s)
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count += 1
		} else if count != 0 {
			return count
		}
	}
	return count
}

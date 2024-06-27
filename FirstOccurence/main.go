package main

import (
	"fmt"
)

func main() {
	result := strStr("mississippi", "issip")
	fmt.Println(result)
}

func strStr(haystack string, needle string) int {
	lengthOfNeedle := len(needle)
	for i := 0; i <= len(haystack)-lengthOfNeedle; i++ {
		if haystack[i:i+lengthOfNeedle] == needle {
			return i
		}
	}
	return -1
}

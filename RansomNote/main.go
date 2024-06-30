package main

import (
	"fmt"
)

func main() {
	result := canConstruct("aab", "baa")
	fmt.Println(result)
}
func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[string]int)

	for _, letter := range ransomNote {
		if _, ok := letters[string(letter)]; ok {
			letters[string(letter)] += 1
		} else {
			letters[string(letter)] = 1
		}
	}

	for _, letter := range magazine {
		if _, ok := letters[string(letter)]; ok {
			letters[string(letter)] -= 1
		}
	}

	for _, v := range letters {
		if v > 0 {
			return false
		}
	}

	return true
}

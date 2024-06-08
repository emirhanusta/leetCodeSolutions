package main

import "fmt"

func main() {
	result := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 {
			if len(strs[i]) >= len(prefix) && strs[i][:len(prefix)] == prefix {
				break
			} else {
				prefix = prefix[:len(prefix)-1]
			}
		}
	}

	return prefix
}

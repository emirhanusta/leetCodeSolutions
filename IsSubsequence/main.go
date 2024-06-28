package main

import "fmt"

func main() {
	result := isSubsequence("", "ahbgdc")
	fmt.Println(result)
}
func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t); i++ {
		if j < len(s) && s[j] == t[i] {
			j++
		}
	}
	return j == len(s)
}

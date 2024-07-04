package main

import "fmt"

func main() {
	result := isAnagram("anagram", "nagaram")
	fmt.Println(result)
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)

	for i := range s {
		m[s[i]]++
		m[t[i]]--
	}

	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}

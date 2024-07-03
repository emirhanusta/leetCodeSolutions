package main

import (
	"fmt"
	"strings"
)

func main() {
	result := wordPattern("jquery", "jquery")
	fmt.Println(result)
}
func wordPattern(pattern string, s string) bool {
	mPattern := make(map[string]string)
	mS := make(map[string]string)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i, item := range pattern {
		if _, ok := mPattern[string(item)]; !ok {
			mPattern[string(item)] = words[i]
		} else {
			if mPattern[string(item)] != words[i] {
				return false
			}
		}
		if _, ok := mS[words[i]]; !ok {
			mS[words[i]] = string(item)
		} else {
			if mS[words[i]] != string(item) {
				return false
			}
		}
	}

	return true
}

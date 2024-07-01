package main

import "fmt"

func main() {
	result := isIsomorphic("badc", "baba")
	fmt.Println(result)
}

func isIsomorphic(s string, t string) bool {
	sMap := make(map[string]string)
	tMap := make(map[string]string)

	for i := 0; i < len(s); i++ {
		if _, ok := sMap[string(s[i])]; ok {
			if sMap[string(s[i])] != string(t[i]) {
				return false
			}
		} else {
			sMap[string(s[i])] = string(t[i])
		}
		if _, ok := tMap[string(t[i])]; ok {
			if tMap[string(t[i])] != string(s[i]) {
				return false
			}
		} else {
			tMap[string(t[i])] = string(s[i])
		}
	}

	return true
}

package main

import "fmt"

func main() {
	result := isValid("[[")
	fmt.Println(result)
}
func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}
	stack := make([]rune, 0)

	for _, value := range s {
		if c, ok := m[value]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

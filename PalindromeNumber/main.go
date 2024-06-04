package main

import (
	"fmt"
)

func main() {
	result := isPalindrome(121)
	fmt.Println(result)
}

func isPalindrome(x int) bool {

	/* string conversion

	xString := strconv.Itoa(x)

	for i := 0; i < len(xString)-1; i++ {
		if string(xString[i]) != string(xString[len(xString)-i-1]) {
			return false
		}
	}*/

	if x < 0 {
		return false
	}
	num := x
	reversed := 0
	for num != 0 {
		reversed = 10*reversed + num%10
		num /= 10
	}
	return reversed == x
}

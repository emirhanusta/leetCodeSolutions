package main

func main() {
	result := romanToInt("MCMXCIV")
	println(result)
}

func romanToInt(s string) int {
	romans := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	value := 0
	if len(s) == 1 {
		return romans[s[0]]
	}
	for i := 0; i < len(s)-1; i++ {
		if romans[s[i]] >= romans[s[i+1]] {
			value += romans[s[i]]
		} else {
			value += romans[s[i+1]] - romans[s[i]]
			i++
		}
		if i == len(s)-2 {
			value += romans[s[i+1]]
		}
	}
	return value
}

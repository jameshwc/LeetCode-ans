package main

import "fmt"

var roman2Number = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			switch s[i] {
			case 'I':
				switch s[i+1] {
				case 'X', 'V':
					ans--
					continue
				}
			case 'X':
				switch s[i+1] {
				case 'L', 'C':
					ans -= 10
					continue
				}
			case 'C':
				switch s[i+1] {
				case 'D', 'M':
					ans -= 100
					continue
				}
			}
		}
		ans += roman2Number[s[i]]
	}
	return ans
}

func romanToIntRevisit(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		nextValue := 0
		if i+1 < len(s) {
			nextValue = roman2Number[s[i+1]]
		}
		v := roman2Number[s[i]]
		if nextValue > v {
			ans -= v
			continue
		}
		ans += v
	}
	return ans
}

func main() {
	fmt.Println(romanToIntRevisit("MCMXCIV"))
}

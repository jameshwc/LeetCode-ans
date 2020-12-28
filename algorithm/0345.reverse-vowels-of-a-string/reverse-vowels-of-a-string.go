package main

import "fmt"

func reverseVowels(s string) string {
	var vowel []byte
	for i := range s {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel = append(vowel, s[i])
		}
	}
	cnt := len(vowel) - 1
	new := make([]byte, len(s))
	for i := range s {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			new[i] = vowel[cnt]
			cnt--
		default:
			new[i] = s[i]
		}
	}
	return string(new)
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

package main

import "fmt"

func findAnagrams(s string, p string) []int {
	counter := make([]int, 26)
	for i := range p {
		counter[p[i]-'a']++
	}
	var ans []int
	check := make([]int, 26)
	for i := range s {
		check[s[i]-'a']++
		if i+1 < len(p) {
			continue
		}
		isMatch := true
		for j := 0; j < 26; j++ {
			if check[j] != counter[j] {
				isMatch = false
				break
			}
		}
		if isMatch {
			ans = append(ans, i-len(p)+1)
		}
		check[s[i-len(p)+1]-'a']--
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

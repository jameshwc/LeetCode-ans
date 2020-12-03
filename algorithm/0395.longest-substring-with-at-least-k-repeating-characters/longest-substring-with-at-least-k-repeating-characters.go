package main

import "fmt"

func longestSubstring(s string, k int) int {
	bucket := make([]int, 26)
	for i := range s {
		bucket[s[i]-'a']++
	}
	left := 0
	ans := 0
	ok := true
	for i := range s {
		if bucket[s[i]-'a'] < k {
			ok = false
			if v := longestSubstring(s[left:i], k); v > ans {
				ans = v
			}
			left = i + 1
		}
	}
	if ok {
		return len(s)
	}
	if v := longestSubstring(s[left:], k); v > ans {
		return v
	}
	return ans
}

func main() {
	fmt.Println(longestSubstring("bbaaacbd", 3))
}

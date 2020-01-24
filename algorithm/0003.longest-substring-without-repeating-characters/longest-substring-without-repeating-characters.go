package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var idx [128]int
	var ans, left int
	for i, v := range s {
		if ans < i-left {
			ans = i - left
		}
		if nowCharID := idx[v]; nowCharID > left {
			left = nowCharID
		}
		idx[v] = i + 1
	}
	if len(s)-left > ans {
		return len(s) - left
	}
	return ans
}

func main() {
	s := "abcabcbb"
	t := "bbbbb"
	m := "pwwkew"
	n := "aab"
	j := "jbpnbww"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring(t))
	fmt.Println(lengthOfLongestSubstring(m))
	fmt.Println(lengthOfLongestSubstring(n))
	fmt.Println(lengthOfLongestSubstring(j))

}

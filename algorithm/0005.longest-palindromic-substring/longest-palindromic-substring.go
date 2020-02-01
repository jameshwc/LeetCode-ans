package main

import (
	"fmt"
	"strings"
)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func preProcess(s string) string {
	ss := []string(strings.Split(s, ""))
	return "#" + strings.Join(ss, "#") + "#"
}

// Manacher's algorithm, solved in linear time complexity.
// details in https://segmentfault.com/a/1190000003914228
func manacher(s string) string {
	t := preProcess(s)
	pos, maxLen, maxRight, n := 0, 0, 0, len(t)
	P := make([]int, n)
	for i := 0; i < n; i++ {
		if i < maxRight {
			P[i] = min(P[pos-(i-pos)], maxRight-i)
		} else {
			P[i] = 1
		}
		for i-P[i] >= 0 && i+P[i] < n && t[i-P[i]] == t[i+P[i]] {
			P[i]++
		}
		if P[i]+i-1 > maxRight {
			maxRight = P[i] + i - 1
			pos = i
		}
	}
	for i := 1; i < n-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i] - 1
			pos = i
		}
	}
	start := (pos - maxLen) / 2
	return s[start : start+maxLen]
}
func checkPalindrome(s string, start int, end int, l int, maxLen int, maxP string) (int, string) {
	for start >= 0 && end < l && s[start] == s[end] {
		start--
		end++
	}
	pLen := end - start - 1
	if pLen > maxLen {
		maxP = s[start+1 : end]
		return pLen, maxP
	} else {
		return maxLen, maxP
	}
}

// The time complexity of this solution is n^2.
func squareLongestPalindrome(s string) string {
	var maxLen = 0
	var maxP string
	if len(s) == 1 {
		return s
	}
	for i, l := 0, len(s); i < l-1; i++ {
		maxLen, maxP = checkPalindrome(s, i, i, l, maxLen, maxP)
		maxLen, maxP = checkPalindrome(s, i, i+1, l, maxLen, maxP)
	}
	return maxP
}

func isPalindrome(s string) bool {
	var j = len(s)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[j-i-1] {
			return false
		}
	}
	return true
}
func cubicLongestPalindrome(s string) string {
	var maxLen = 0
	var maxP string
	var l = len(s)
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			if isPalindrome(s[j : i+1]) {
				if i-j+1 > maxLen {
					maxLen = i - j + 1
					maxP = s[j : i+1]
				}
			}
		}
	}
	return maxP
}

func main() {
	fmt.Println(manacher("babad"))
	fmt.Println(manacher("cbbd"))
	// if longestPalindrome("babad") != "bab" || longestPalindrome("cbbd") != "bb" {
	// 	fmt.Println("Wrong Answer!")
	// } else {
	// 	fmt.Println("Accepted!")
	// }
}

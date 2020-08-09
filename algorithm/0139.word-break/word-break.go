package main

import "fmt"

func find(s, cur string, wordDict []string) bool {
	if s == cur {
		return true
	}
	for i := range wordDict {
		if len(cur)+len(wordDict[i]) <= len(s) && s[len(cur):len(cur)+len(wordDict[i])] == wordDict[i] && find(s, cur+wordDict[i], wordDict) {
			return true
		}
	}
	return false
}

// TLE
func wordBreak_Recursion(s string, wordDict []string) bool {
	isFind := false
	var find func(string)
	find = func(cur string) {
		if s == cur {
			isFind = true
		}
		if isFind == true {
			return
		}
		for i := range wordDict {
			if len(cur)+len(wordDict[i]) <= len(s) && s[len(cur):len(cur)+len(wordDict[i])] == wordDict[i] {
				find(cur + wordDict[i])
			}
			if isFind == true {
				return
			}
		}
	}
	find("")
	return isFind
}

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		set[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := set[s[j:i]]; ok {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)]
}
func main() {
	// fmt.Println(wordBreak("bb", []string{"b"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}) && wordBreak("applepenapple", []string{"apple", "pen"}) && !wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

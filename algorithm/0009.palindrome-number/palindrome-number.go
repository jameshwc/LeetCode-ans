package main

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func isPalindromeRevisit(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revNum := 0
	for x > revNum {
		revNum = revNum*10 + x%10
		x /= 10
	}
	if x == revNum || x == revNum/10 {
		return true
	}
	return false
}

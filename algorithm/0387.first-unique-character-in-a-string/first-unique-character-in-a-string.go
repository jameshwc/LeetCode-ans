package main

func firstUniqChar(s string) int {
	bucket := make([]int, 26)
	for i := range s {
		bucket[s[i]-'a']++
	}
	for i := range s {
		if bucket[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

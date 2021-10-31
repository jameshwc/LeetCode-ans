func isPalindrome(s string) bool {
	t := make([]byte, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			t[j] = ch - 'A' + 'a'
			j++
		} else if (ch >= 'a' && ch <= 'z') || (ch <= '9' && ch >= '0') {
			t[j] = ch
			j++
		}
	}
	for i := 0; i < j; i++ {
		if t[i] != t[j-i-1] {
			return false
		}
	}
	return true
}
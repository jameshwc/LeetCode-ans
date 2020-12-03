func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := -1
	same := true
	for same {
		s++
		ch := byte(0)
		for i := range strs {
			if len(strs[i]) == s {
				same = false
				break
			}
			if i == 0 {
				ch = strs[i][s]
			}
			if ch != strs[i][s] {
				same = false
				break
			}
		}
	}
	return strs[0][:s]
}
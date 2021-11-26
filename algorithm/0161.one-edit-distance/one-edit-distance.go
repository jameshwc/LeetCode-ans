func isSameString(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

func isOneEditDistance(s string, t string) bool {

	maxLen := len(t)
	if maxLen < len(s) {
		maxLen = len(s)
	}

	for i := 0; i < maxLen; i++ {
		if i == len(s) {
			return len(t)-len(s) <= 1
		} else if i == len(t) {
			return len(s)-len(t) <= 1
		}

		if s[i] != t[i] {
			return isSameString(s[i+1:], t[i:]) || // delete s[i]
				isSameString(s[i:], t[i+1:]) || // insert s[i-1]
				isSameString(s[i+1:], t[i+1:]) // replace s[i]
		}
	}
	return false
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func editDistance(s string, t string) int {

	l1, l2 := len(s), len(t)
	dp := make([][]int, l1+1)

	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
			}
		}
	}
	return dp[l1][l2]
}
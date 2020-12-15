func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	cnt := 0
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s)-1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				cnt++
			}
		}
	}
	return cnt
}

// abccab
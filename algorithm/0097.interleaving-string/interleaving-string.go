package algorithm

func bruteForce(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if len(s1) == 0 {
		return s2 == s3
	}

	if len(s2) == 0 {
		return s1 == s3
	}

	if len(s3) == 0 {
		return false
	}

	if s1[0] != s3[0] && s2[0] != s3[0] {
		return false
	}

	if s1[0] == s3[0] {
		if s2[0] == s3[0] {
			return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
		}
		return isInterleave(s1[1:], s2, s3[1:])
	}

	return isInterleave(s1, s2[1:], s3[1:])
}

func recursion(s1, s2, s3 string) bool {
	var f func(i, j, k int) bool

	memo := make([][]int, len(s1))
	for i := range memo {
		memo[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			memo[i][j] = -1
		}
	}

	f = func(i, j, k int) bool {
		if i == len(s1) {
			return s2[j:] == s3[k:]
		}

		if j == len(s2) {
			return s1[i:] == s3[k:]
		}

		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		ans := false
		if s1[i] == s3[k] && f(i+1, j, k+1) {
			ans = true
		} else if s2[j] == s3[k] && f(i, j+1, k+1) {
			ans = true
		}

		if ans {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}

		return ans
	}

	return f(0, 0, 0)
}

func twoDArray(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j] == s3[j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i] == s3[i]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i] == s3[i+j]) || (dp[i][j-1] && s2[j] == s3[i+j])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func isInterleave(s1, s2, s3 string) bool {

}

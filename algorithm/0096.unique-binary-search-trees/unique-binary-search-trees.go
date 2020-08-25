func numTrees(n int) int {
	dp := make([]int, 20) // n's max = 19
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	// dp[3] = dp[0] * dp[2] + dp[1] * dp[1] + dp[2] * dp[0]
	// dp[4] = dp[0] * dp[3] + dp[2] * dp[1] + dp[1] * dp[2] + dp[0] * dp[3]
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
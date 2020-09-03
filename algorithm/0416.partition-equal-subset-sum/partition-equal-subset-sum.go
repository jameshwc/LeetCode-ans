func canPartition(nums []int) bool {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, total+1)
	}
	dp[0][nums[0]] = true
	dp[0][0] = true
	for i := 1; i < len(nums); i++ {
		dp[i][nums[i]] = true
		for j := 0; j <= total; j++ {
			if dp[i-1][j] || (j-nums[i] >= 0 && dp[i-1][j-nums[i]]) {
				dp[i][j] = true
			}
		}
	}
	return dp[len(nums)-1][target]
}
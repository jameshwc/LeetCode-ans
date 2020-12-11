func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = 1 << 31
	}
	dp[0] = 0
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := range coins {
		m := coins[i]
		for j := 0; j+m <= amount; j++ {
			dp[j+m] = min(dp[j+m], dp[j]+1)
		}
	}
	if dp[amount] == 1<<31 {
		return -1
	}
	return dp[amount]
}
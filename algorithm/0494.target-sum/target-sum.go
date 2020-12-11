package main

// func findTargetSumWays(nums []int, S int) int {
// 	cnt := 0
// 	var bruteForce func(int, int)
// 	bruteForce = func(sum, idx int) {
// 		if idx == len(nums) {
// 			if sum == S {
// 				cnt++
// 			}
// 			return
// 		}
// 		bruteForce(sum+nums[idx], idx+1)
// 		bruteForce(sum-nums[idx], idx+1)
// 	}
// 	bruteForce(0, 0)
// 	return cnt
// }

func findTargetSumWays(nums []int, S int) int {
	if S > 1000 {
		return 0
	}
	dp := make([]int, 2001)
	dp[1000-nums[0]]++
	dp[1000+nums[0]]++
	for i := 1; i < len(nums); i++ {
		next := make([]int, 2001)
		for j := 0; j <= 2000; j++ {
			if dp[j] > 0 {
				next[j+nums[i]] += dp[j]
				next[j-nums[i]] += dp[j]
			}
		}
		dp = next
	}
	return dp[1000+S]
}

func main() {
	findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1)
}

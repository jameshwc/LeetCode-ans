package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1 << 31
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 0; i+j*j <= n; j++ {
			if dp[i]+1 < dp[i+j*j] {
				dp[i+j*j] = dp[i] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12), numSquares(13))
}

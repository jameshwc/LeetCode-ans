package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func maxUncrossedLines(A []int, B []int) int {
	dp := make([][]int, len(A))
	for i := range dp {
		dp[i] = make([]int, len(B))
	}
	for i := range A {
		for j := range B {
			if A[i] == B[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i > 0 {
					dp[i][j] = max(dp[i-1][j], dp[i][j])
				}
				if j > 0 {
					dp[i][j] = max(dp[i][j-1], dp[i][j])
				}
			}
		}
	}
	return dp[len(A)-1][len(B)-1]
}

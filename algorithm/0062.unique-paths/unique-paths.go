package main

import "fmt"

func uniquePaths(m int, n int) int {
	var dp [100][100]int
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	if uniquePaths(7, 3) == 28 && uniquePaths(3, 2) == 3 {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}

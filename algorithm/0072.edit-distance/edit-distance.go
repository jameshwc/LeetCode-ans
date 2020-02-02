package main

import "fmt"

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
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
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[l1][l2]
}

func main() {
	if minDistance("horse", "ros") != 3 || minDistance("intention", "execution") != 5 {
		fmt.Println("Wrong Answer!")
	} else {
		fmt.Println("Accepted!")
	}
}

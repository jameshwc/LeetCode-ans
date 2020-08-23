package main

import "fmt"

/* Brute-Force */
func maximalSquareBruteForce(matrix [][]byte) int {
	maxArea := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				continue
			}
			endi, endj := i, j
			for {
				endi++
				endj++
				if endi == len(matrix) || endj == len(matrix[endi]) || matrix[endi][endj] == '0' {
					break
				}
				cur := i
				for ; cur < endi; cur++ {
					if matrix[cur][endj] == '0' {
						break
					}
				}
				if matrix[cur][endj] == '0' {
					break
				}
				cur = j
				for ; cur < endj; cur++ {
					if matrix[endi][cur] == '0' {
						break
					}
				}
				if matrix[endi][cur] == '0' {
					break
				}
			}
			if area := (endi - i) * (endi - i); maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

/* Dynamic Programming */
func maximalSquareDP(matrix [][]byte) int {
	maxL := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i])+1)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			if maxL < dp[i][j] {
				maxL = dp[i][j]
			}
		}
	}
	return maxL * maxL
}

/* Dynamic Programming inplace */
func maximalSquare(matrix [][]byte) int {
	maxL := 0
	// dp := make([]int, len(matrix))
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] -= '0'
			if i != 0 && j != 0 && matrix[i][j] == 1 {
				matrix[i][j] = byte(min(int(matrix[i-1][j]), int(matrix[i][j-1]), int(matrix[i-1][j-1])) + int(matrix[i][j]))
			}
			if L := int(matrix[i][j]); L > maxL {
				maxL = L
			}
		}
	}
	return maxL * maxL
}

/* Dynamic Programming with rolling array (need to revisit) */
func maximalSquare(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}
	dp := make([]int, len(matrix[0])+1)
	maxL := 0
	for i := 0; i < len(matrix); i++ {
		prev := 0
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				temp := dp[j+1]
				dp[j+1] = min(prev, dp[j+1], dp[j]) + 1
				if dp[j+1] > maxL {
					maxL = dp[j+1]
				}
				prev = temp
			} else {
				dp[j+1] = 0
			}
		}
	}
	return maxL * maxL
}

func main() {
	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'},
	// }
	matrix := [][]byte{
		{'0', '0', '0', '1', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0'},
		{'0', '1', '0', '1', '0', '0', '1'},
		{'0', '0', '1', '1', '0', '0', '1'},
		{'1', '1', '1', '1', '1', '1', '0'},
		{'1', '0', '0', '1', '0', '1', '1'},
		{'0', '1', '0', '0', '1', '0', '1'},
		{'1', '1', '0', '1', '1', '1', '0'},
		{'1', '0', '1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '0', '0', '0'},
	}
	fmt.Println(maximalSquareDP(matrix))
}

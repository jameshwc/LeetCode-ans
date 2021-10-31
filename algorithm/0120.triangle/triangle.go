func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
			} else if j == i {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}
	ans := 1 << 31
	for i := 0; i < len(triangle); i++ {
		ans = min(ans, triangle[len(triangle)-1][i])
	}
	return ans
}
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		for j := range ans[i] {
			if j == 0 || j == i {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}

	return ans
}
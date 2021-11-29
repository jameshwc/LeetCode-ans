func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2+n%2; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]

		}
	}
}
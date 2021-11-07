func setZeroes(matrix [][]int) {

	var rowFlag, colFlag bool

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {

				if i == 0 {
					rowFlag = true
				}

				if j == 0 {
					colFlag = true
				}

				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowFlag {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if colFlag {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		check := func(r int, rev bool) bool {
			rec := make([]bool, 10)
			for j := 0; j < 9; j++ {
				ch := board[i][j]
				if rev {
					ch = board[j][i]
				}
				if ch != '.' {
					if rec[ch-'0'] == true {
						return false
					}
					rec[ch-'0'] = true
				}
			}
			return true
		}
		if !check(i, true) || !check(i, false) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rec := make([]bool, 10)
			check := func(r, c int) bool {
				ch := board[i*3+r][j*3+c]
				if ch != '.' {
					if rec[ch-'0'] == true {
						return false
					}
					rec[ch-'0'] = true
				}
				return true
			}
			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					if !check(r, c) {
						return false
					}
				}
			}
		}
	}
	return true
}
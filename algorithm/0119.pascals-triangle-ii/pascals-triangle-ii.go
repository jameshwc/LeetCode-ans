func getRowCombination(rowIndex int) []int {
	C := func(upper, lower int) int {
		if lower > upper-lower {
			lower = upper - lower
		}
		result := 1
		for d := 1; d <= lower; d++ {
			result *= upper
			upper--
			result /= d
		}
		return result
	}
	ans := make([]int, rowIndex+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = int(C(rowIndex, i))
		ans[len(ans)-i-1] = ans[i]
	}
	return ans
}

func getRow(rowIndex int) []int {
	var prev []int
	var ans []int
	for i := 0; i <= rowIndex; i++ {
		ans = make([]int, i+1)
		for j := range ans {
			if j == 0 || j == i {
				ans[j] = 1
			} else {
				ans[j] = prev[j-1] + prev[j]
			}
		}
		prev = ans
	}

	return ans
}
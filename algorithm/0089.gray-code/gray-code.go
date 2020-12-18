package main

func grayCode(n int) []int {
	var ans = []int{0}
	for i := 0; i < n; i++ {
		add := 1 << i
		for g := len(ans) - 1; g >= 0; g-- {
			ans = append(ans, ans[g]+add)
		}
	}
	return ans
}

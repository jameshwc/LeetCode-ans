func countBits(num int) []int {
	var ans = make([]int, num+1)
	ans[0] = 0
	for i := 1; i <= num; i++ {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}
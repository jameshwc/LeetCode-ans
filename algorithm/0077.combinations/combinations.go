func combine(n int, k int) [][]int {
	var ans [][]int

	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	var dfs func(int, []int)
	var start []int
	dfs = func(index int, cur []int) {

		if len(cur) == k {
			cc := make([]int, k)
			copy(cc, cur)
			ans = append(ans, cc)
			return
		}

		if index == len(nums) {
			return
		}

		dfs(index+1, cur)
		cur = append(cur, nums[index])
		dfs(index+1, cur)
	}

	dfs(0, start)
	return ans
}
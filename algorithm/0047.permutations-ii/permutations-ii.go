func permuteUnique(nums []int) [][]int {
	var result [][]int
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(nums) {
			result = append(result, append([]int{}, nums...))
			return
		}
		used := make(map[int]struct{})
		for i := start; i < len(nums); i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			used[nums[i]] = struct{}{}
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	backtrack(0)
	return result
}

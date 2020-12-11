func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
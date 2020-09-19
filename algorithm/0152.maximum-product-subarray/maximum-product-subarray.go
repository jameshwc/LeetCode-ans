func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxProduct(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	neg := nums[0]
	pos := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		negTemp := neg
		neg = min(neg * nums[i], min(pos * nums[i], nums[i]))
		pos[i] = max(negTemp * nums[i], max(pos * nums[i], nums[i]))
		ans = max(ans, pos[i])
	}
	return ans
}
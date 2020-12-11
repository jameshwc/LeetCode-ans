func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	max := 0
	for i := range nums {
		if nums[i] == 1 {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	if cnt > max {
		max = cnt
	}
	return max
}
func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := 1 << 31
	threeSum := make([]int, 3)
	for i := 0; i+1 < len(nums); i++ {
		rem := target - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left != right {
			cur := rem - nums[left] - nums[right]
			if abs(cur) < min {
				min = abs(cur)
				threeSum[0] = nums[i]
				threeSum[1] = nums[left]
				threeSum[2] = nums[right]
			}
			if cur > 0 {
				left++
			} else if cur < 0 {
				right--
			} else {
				return target
			}
		}
	}
	return threeSum[0] + threeSum[1] + threeSum[2]
}
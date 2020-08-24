func productExceptSelf(nums []int) []int {
	prefix := 1
	output := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			output[i] = nums[i]
		} else {
			output[i] = output[i+1] * nums[i]
		}
	}
	for i := range nums {
		if i == 0 {
			output[i] = output[i+1]
			prefix = nums[i]
		} else if i+1 == len(nums) {
			output[i] = prefix
		} else {
			output[i] = prefix * output[i+1]
			prefix *= nums[i]
		}
	}
	return output
}
package sorting

func lowerBound(nums []int, left, right, target int) int {
	for left < right {
		mid := left + ((right - left) >> 1)
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func upperBound(nums []int, left, right, target int) int {
	for left < right {
		mid := left + ((right - left) >> 1)
		if target >= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

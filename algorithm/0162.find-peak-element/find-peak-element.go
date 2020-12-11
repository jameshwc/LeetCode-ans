package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if len(nums) > 1 && nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

func findPeakElementRevisit(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left, right int) int {
	if left == right {
		return left
	}
	mid := (left + right) >> 1
	if nums[mid] < nums[mid+1] {
		return search(nums, mid+1, right)
	}
	return search(nums, left, mid)
}

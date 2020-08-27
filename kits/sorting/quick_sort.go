package sorting

func QuickSort(nums []int, left, right int) {
	if right < left {
		return
	}
	pivot := nums[right]
	cur := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[cur] = nums[cur], nums[i]
			cur++
		}
	}
	nums[right], nums[cur] = nums[cur], nums[right]
	QuickSort(nums, left, cur-1)
	QuickSort(nums, cur+1, right)
}

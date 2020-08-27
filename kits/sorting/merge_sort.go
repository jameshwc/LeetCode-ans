package sorting

func merge(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	merge(nums[:n/2])
	merge(nums[n/2:])
	i, j := 0, n/2
	cur := 0
	sortNum := make([]int, len(nums))
	for {
		if nums[i] > nums[j] {
			sortNum[cur] = nums[j]
			j++
			cur++
		} else {
			sortNum[cur] = nums[i]
			i++
			cur++
		}
		if i == n/2 {
			for ; j < len(nums); j++ {
				sortNum[cur] = nums[j]
				cur++
			}
			break
		}
		if j == len(nums) {
			for ; i < n/2; i++ {
				sortNum[cur] = nums[i]
				cur++
			}
			break
		}
	}
	for i := range nums {
		nums[i] = sortNum[i]
	}
}
func MergeSort(nums []int) {
	merge(nums)
}

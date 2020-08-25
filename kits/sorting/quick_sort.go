package sorting

import "fmt"

func quickSort(nums []int, left, right int) {
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
	quickSort(nums, left, cur-1)
	quickSort(nums, cur+1, right)
}

func main() {
	nums := []int{1, 7, 2, 3, 5, 10, 8, 9, 12, 3, 2, 4}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1}
	fmt.Println(searchInsert(nums, 0))
}

package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		if mid+1 == len(nums) || mid == 0 {
			return nums[mid]
		}

		if nums[mid] == nums[mid-1] {
			if mid%2 == 0 {
				right = mid
			} else {
				left = mid + 1
			}
		} else if nums[mid] == nums[mid+1] {
			if mid%2 == 0 {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			return nums[mid]
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))
}

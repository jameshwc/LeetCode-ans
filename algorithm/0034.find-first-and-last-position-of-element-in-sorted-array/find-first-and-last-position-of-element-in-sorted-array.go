package main

import "fmt"

// TODO: how to one-pass it?
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	search := func(isLeft bool) int {
		left, right := 0, n-1
		for left != right {
			mid := (left + right) / 2
			if nums[mid] > target || (nums[mid] == target && isLeft) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if !isLeft && nums[left] != target && left > 0 && nums[left-1] == target {
			return left - 1
		} else if nums[left] != target {
			return -1
		}
		return right
	}
	return []int{search(true), search(false)}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10, 10, 10}, 10))
}

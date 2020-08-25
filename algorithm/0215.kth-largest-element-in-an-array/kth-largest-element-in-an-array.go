package main

import "fmt"

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	cur := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[cur] = nums[cur], nums[i]
			cur++
		}
	}
	nums[cur], nums[right] = nums[right], nums[cur]
	return cur
}

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	k = len(nums) - k // lazy sol because partition returns kth smallest num idx
	for {
		pos := partition(nums, left, right)
		if pos == k {
			return nums[pos]
		} else if pos > k {
			right = pos - 1
		} else {
			left = pos + 1
		}
	}
}

func main() {
	if findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5 && findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) == 4 {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}

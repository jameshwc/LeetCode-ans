package main

import "fmt"

func maxSubArray(nums []int) int {
	max, cur := nums[0], 0
	for i := range nums {
		if cur+nums[i] > nums[i] {
			cur = cur + nums[i]
		} else {
			cur = nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

func main() {
	if maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6 {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}

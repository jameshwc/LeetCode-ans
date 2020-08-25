package main

import "fmt"

func DPlengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	maxV := 0
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if maxV < dp[i] {
			maxV = dp[i]
		}
	}
	return maxV
}

func lowerBound(nums []int, n, target int) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) >> 1
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	id := 0
	for i := range nums {
		if i == 0 {
			dp[id] = nums[i]
		} else {
			if nums[i] > dp[id] {
				id++
				dp[id] = nums[i]
			} else {
				idx := lowerBound(dp, id+1, nums[i])
				dp[idx] = nums[i]
			}
		}
		continue
	}
	return id + 1
}
func main() {
	nums := []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}
	fmt.Println(lengthOfLIS(nums))
}

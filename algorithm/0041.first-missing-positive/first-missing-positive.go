package main

import "fmt"

func boolArray(nums []int) int {
	var appear [1 << 20]bool
	n := len(nums)
	if n == 0 {
		return 1
	}
	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] < 1 {
			continue
		} else {
			appear[nums[i]] = true
		}
	}
	for i := 1; i <= n; i++ {
		if appear[i] == false {
			return i
		}
	}
	return n + 1
}
func firstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] != i+1 && nums[i] > 0 && nums[i] < n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	if firstMissingPositive([]int{1, 2, 0}) != 3 ||
		firstMissingPositive([]int{3, 4, -1, 1}) != 2 ||
		firstMissingPositive([]int{7, 8, 9, 11, 12}) != 1 {
		fmt.Println("Wrong Answer!")
	} else {
		fmt.Println("Accepted!")
	}
}

package main

import "fmt"

/* complicate solution */
func MymoveZeroes(nums []int) {
	leftZeroIdx := -1
	for i := range nums {
		if nums[i] == 0 && leftZeroIdx == -1 {
			j := i + 1
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j != len(nums) {
				nums[i], nums[j] = nums[j], nums[i]
				leftZeroIdx = i + 1
			}
		} else if leftZeroIdx != -1 && nums[i] != 0 {
			nums[leftZeroIdx], nums[i] = nums[i], nums[leftZeroIdx]
			leftZeroIdx++
		}
	}
}

/* simple solution but hard to think of */
func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{0, 0, 1, 2, 5}
	moveZeroes(nums)
	fmt.Println(nums)
}

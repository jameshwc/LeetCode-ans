package main

import "fmt"

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		switch nums[i] {
		case 0:
			nums[i], nums[left] = nums[left], nums[i]
			left++
		case 2:
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
}
func countingSortColors(nums []int) {
	var counter [3]int
	for i := range nums {
		counter[nums[i]]++
	}
	for i, idx := 0, 0; i < 3; i++ {
		for j := 0; j < counter[i]; j++ {
			nums[idx] = i
			idx++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	// sortColors(nums)
	countingSortColors(nums)
	fmt.Println(nums)
}

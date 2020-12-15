package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	mx, mn := -(1 << 31), 1<<31
	begin, end := -100, -101
	for i := 0; i < len(nums); i++ {
		if nums[i] >= mx {
			mx = nums[i]
		} else {
			end = i
		}
		if nums[len(nums)-i-1] <= mn {
			mn = nums[len(nums)-i-1]
		} else {
			begin = len(nums) - i - 1
		}
	}
	return end - begin + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 2, 2}))
	fmt.Println(findUnsortedSubarray([]int{3, 10, 5, 2, 0, 1}))

}

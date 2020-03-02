package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
func canJump(nums []int) bool {
	n := len(nums)
	start, reach := 0, 0
	for start <= reach {
		reach = max(start+nums[start], reach)
		if reach >= n-1 {
			return true
		}
		start++
	}
	return false
}

func main() {
	if canJump([]int{2, 3, 1, 1, 4}) && !canJump([]int{3, 2, 1, 0, 4}) && canJump([]int{0}) {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Wrong Answer")
	}
}

package main

import "fmt"

func singleNumber(nums []int) int {
	var ans int
	for i := range nums {
		ans = ans ^ nums[i]
	}
	return ans
}

func main() {
	if singleNumber([]int{2, 2, 1}) != 1 || singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		fmt.Println("Wrong Answer")
	} else {
		fmt.Println("Accepted")
	}
}

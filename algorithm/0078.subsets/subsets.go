package main

import (
	"fmt"
)

func recurGenSubset(ans *[][]int, nums []int, cur []int, id int, n int) {
	if id == n {
		cp := make([]int, len(cur))
		copy(cp, cur) // should avoid shallow copy here!
		*ans = append(*ans, cp)
		return
	}
	recurGenSubset(ans, nums, cur, id+1, n)
	cur = append(cur, nums[id])
	recurGenSubset(ans, nums, cur, id+1, n)
}
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	cur := make([]int, 0)
	recurGenSubset(&ans, nums, cur, 0, n)
	return ans
}

func main() {
	nums := []int{0, 3, 5, 7, 9}
	fmt.Println(subsets(nums))
}

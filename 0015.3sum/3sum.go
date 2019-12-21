package problem0015

import (
	"sort"
)

// Runtime: 1652 ms, faster than 9.37% of Go online submissions for 3Sum.
// Memory Usage: 272.9 MB, less than 41.67% of Go online submissions for 3Sum.
type tuple struct {
	firstNum,
	secondNum,
	thirdNum int
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var check = make(map[tuple](bool))
	l := len(nums)
	if l < 3 {
		return nil
	}
	for i := 0; i < l; i++ {
		var index = make(map[int](int))
		target := -nums[i]
		for j := i + 1; j < l; j++ {
			if n, ok := index[target-nums[j]]; ok {
				c := tuple{nums[i], nums[n], nums[j]}
				if _, ok := check[c]; !ok {
					ans = append(ans, []int{nums[i], nums[n], nums[j]})
					check[c] = true
				}
			}
			index[nums[j]] = j
		}
	}
	return ans
}

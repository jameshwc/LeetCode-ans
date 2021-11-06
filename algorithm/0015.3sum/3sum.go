package problem0015

import (
	"sort"
)

/*
Runtime: 1480 ms, faster than 10.62% of Go online submissions for 3Sum.
Memory Usage: 205.5 MB, less than 58.33% of Go online submissions for 3Sum.
*/
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

func threeSumSorting(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i+1 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if t := nums[i] + nums[left] + nums[right]; t > 0 {
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				right--
			} else if t < 0 {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return ans
}

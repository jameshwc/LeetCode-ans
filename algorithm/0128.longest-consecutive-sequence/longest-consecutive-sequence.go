package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}
	maxCount := 0
	for k := range set {
		findConsecutive := func(key, add int) int {
			cnt := 0
			for {
				if _, ok := set[key+add]; ok {
					cnt++
					key += add
					delete(set, key)
				} else {
					break
				}
			}
			return cnt
		}
		if count := 1 + findConsecutive(k, 1) + findConsecutive(k, -1); count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func main() {
	input := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(input) == 4)
}

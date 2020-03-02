package main

import "fmt"

func fact(n int) int {
	sum := 1
	for i := 2; i <= n; i++ {
		sum = sum * i
	}
	return sum
}
func reverse(nums []int, begin int, end int) {
	for i, j := begin, end-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int, n int) {
	i, j := n-2, n-1
	for i >= 0 && nums[i] > nums[i+1] {
		i--
	}
	if i >= 0 {
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, n)
}

func permute(nums []int) [][]int {
	t := fact(len(nums))
	n := len(nums)
	var ans [][]int
	for i := 1; i < t; i++ {
		cp := make([]int, n)
		copy(cp, nums)
		ans = append(ans, cp)
		nextPermutation(nums, n)
	}
	ans = append(ans, nums)
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

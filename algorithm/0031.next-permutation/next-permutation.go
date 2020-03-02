package main

func reverse(nums []int, begin int, end int) {
	for i, j := begin, end-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
func nextPermutation(nums []int) {
	n := len(nums)
	i, j := n-2, n-1
	for i >= 0 && nums[i] >= nums[i+1] {
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

func main() {
	input1 := []int{1, 3, 2}
	input2 := []int{2, 3, 1}
	input3 := []int{1, 5, 1}
	nextPermutation(input1)
	nextPermutation(input2)
	nextPermutation(input3)
}

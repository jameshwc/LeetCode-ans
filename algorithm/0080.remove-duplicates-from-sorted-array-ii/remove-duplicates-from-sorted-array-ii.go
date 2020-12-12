package main

func removeDuplicates(nums []int) int {
	cur := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				continue
			}
		} else {
			count = 1
		}
		nums[cur] = nums[i]
		cur++
	}
	return cur
}

func removeDuplicatesRevisit(nums []int) int {
	i := 0
	for _, num := range nums {
		if i < 2 || num != nums[i-2] {
			nums[i] = num
			i++
		}
	}
	return i
}

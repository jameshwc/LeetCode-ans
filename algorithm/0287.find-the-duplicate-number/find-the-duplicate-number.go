func findDuplicate(nums []int) int {
	set := make(map[int]int)
	for i := range nums {
		if _, ok := set[nums[i]]; ok {
			return nums[i]
		} else {
			set[nums[i]] = 1
		}
	}
	return 0
}

func tortoise(nums []int) int {
	fast, slow := nums[0], nums[0]
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}
	slow = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}
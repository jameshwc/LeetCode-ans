package problem0001

func twoSum(nums []int, target int) []int {
	var index = make(map[int](int))
	for i, l := 0, len(nums); i < l; i++ {
		if n, ok := index[target-nums[i]]; ok {
			return []int{n, i}
		}
		index[nums[i]] = i
	}
	return nil
}

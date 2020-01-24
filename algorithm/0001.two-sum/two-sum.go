package problem0001

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Two Sum.
Memory Usage: 3.4 MB, less than 46.15% of Go online submissions for Two Sum.
*/

func twoSum(nums []int, target int) []int {
	var index = make(map[int](int), len(nums))
	for i, v := range nums {
		if n, ok := index[target-v]; ok {
			return []int{n, i}
		}
		index[nums[i]] = i
	}
	return nil
}

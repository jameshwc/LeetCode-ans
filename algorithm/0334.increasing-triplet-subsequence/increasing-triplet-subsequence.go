package main

func increasingTriplet(nums []int) bool {
	smallest, secondSmallest := 1<<31, 1<<31
	for i := range nums {
		switch {
		case nums[i] <= smallest:
			smallest = nums[i]
		case nums[i] <= secondSmallest:
			secondSmallest = nums[i]
		default:
			return true
		}
	}
	return false
}

func threeSum(nums []int, target int) [][]int {
	var ans [][]int
	for i := 0; i+1 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if t := target - (nums[i] + nums[left] + nums[right]); t < 0 {
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				right--
			} else if t > 0 {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				right--
			}
		}
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		rem := target - nums[i]
		threeSumAns := threeSum(nums[i+1:], rem)
		for j := range threeSumAns {
			var cur []int
			cur = append(cur, nums[i])
			cur = append(cur, threeSumAns[j]...)
			ans = append(ans, cur)
		}
	}
	return ans
}

/* General solution */

func kSum(nums []int, target int, k int) [][]int {
	if k == 2 {
		return twoSum(nums, target)
	}

	var ans [][]int

	for i := 0; i < len(nums)-k+1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		rem := target - nums[i]
		degradeAns := kSum(nums[i+1:], rem, k-1)
		for j := range degradeAns {
			var cur []int
			cur = append(cur, nums[i])
			cur = append(cur, degradeAns[j]...)
			ans = append(ans, cur)
		}
	}
	return ans
}

func twoSum(nums []int, target int) (ans [][]int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		if rem := target - nums[left] - nums[right]; rem > 0 {
			for left < right && nums[left] == nums[left+1] {
				left++
			}
			left++
		} else if rem < 0 {
			for left < right && nums[right] == nums[right-1] {
				right--
			}
			right--
		} else {
			ans = append(ans, []int{nums[left], nums[right]})
			for left < right && nums[right] == nums[right-1] {
				right--
			}
			right--
		}
	}
	return
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}
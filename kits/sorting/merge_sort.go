package sorting

import "fmt"

func merge(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	merge(nums[:n/2])
	merge(nums[n/2:])
	i, j := 0, n/2
	cur := 0
	sortNum := make([]int, len(nums))
	for {
		if nums[i] > nums[j] {
			sortNum[cur] = nums[j]
			j++
			cur++
		} else {
			sortNum[cur] = nums[i]
			i++
			cur++
		}
		if i == n/2 {
			for ; j < len(nums); j++ {
				sortNum[cur] = nums[j]
				cur++
			}
			break
		}
		if j == len(nums) {
			for ; i < n/2; i++ {
				sortNum[cur] = nums[i]
				cur++
			}
			break
		}
	}
	for i := range nums {
		nums[i] = sortNum[i]
	}
}
func mergeSort(nums []int) {
	merge(nums)
}

func main() {
	nums := []int{1, 7, 2, 3, 5, 10, 8, 9, 12, 3, 2, 4}
	mergeSort(nums)
	fmt.Println(nums)
}

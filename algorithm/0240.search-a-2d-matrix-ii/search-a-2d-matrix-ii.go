package main

import "fmt"

func binarySearch(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		if mid := (left + right) >> 1; nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	if nums[left] == target {
		return true
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix[0])
	for i := range matrix {
		if matrix[i][0] > target {
			return false
		}
		if matrix[i][m-1] < target {
			continue
		}
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func main() {
	// matrix := [][]int{
	// 	{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, 6, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30},
	// }
	matrix := [][]int{
		{-5},
	}
	fmt.Println(searchMatrix(matrix, -5))
}

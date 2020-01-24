package main

import "fmt"

func findMedian(nums []int) float64 {
	l := len(nums)
	if l%2 != 0 {
		return float64(nums[l/2])
	}
	return float64(nums[l/2-1]+nums[l/2]) / 2.0
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 {
		return findMedian(nums2)
	}
	if l2 == 0 {
		return findMedian(nums1)
	}
	if l1 == 1 && l2 == 1 {
		return float64(nums1[0]+nums2[0]) / 2.0
	}
	m1 := findMedian(nums1)
	m2 := findMedian(nums2)
	// if(m1 <= nums2[0] && nums1[l1])
	var p1, p2 []int
	if m1 == m2 {
		return m1
	} else if m1 > m2 {
		p1 = nums1[:l1/2]
		p2 = nums2[l2/2:]
		fmt.Println(p1, p2)
	} else {
		p1 = nums1[l1/2:]
		p2 = nums2[:l2/2]
	}
	return findMedianSortedArrays(p1, p2)
}
func main() {
	// fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	// fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
}

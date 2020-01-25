package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func findKth(k int, nums1 []int, i int, nums2 []int, j int) int {
	if i > j {
		return findKth(k, nums2, j, nums1, i)
	} else if i == 0 {
		return nums2[k-1]
	} else if k == 1 {
		return min(nums1[0], nums2[0])
	} else {
		var pa = min(k/2, i)
		var pb = k - pa
		if nums1[pa-1] < nums2[pb-1] {
			return findKth(k-pa, nums1[pa:], i-pa, nums2, j)
		} else if nums1[pa-1] > nums2[pb-1] {
			return findKth(k-pb, nums1, i, nums2[pb:], j-pb)
		} else {
			return nums1[pa-1]
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2 = len(nums1), len(nums2)
	if (l1+l2)%2 == 1 {
		return float64(findKth((l1+l2)/2+1, nums1, l1, nums2, l2))
	}
	return float64(findKth((l1+l2)/2, nums1, l1, nums2, l2)+findKth((l1+l2)/2+1, nums1, l1, nums2, l2)) / 2
}

func findMedian(nums []int) float64 {
	l := len(nums)
	if l%2 != 0 {
		return float64(nums[l/2])
	}
	return float64(nums[l/2-1]+nums[l/2]) / 2.0
}
func linearFindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2 = len(nums1), len(nums2)
	if l1 == 0 {
		return findMedian(nums2)
	} else if l2 == 0 {
		return findMedian(nums1)
	}
	var i, j = 0, 0
	var nums = make([]int, l1+l2)
	var cnt = 0
	for i != l1 && j != l2 {
		if nums1[i] < nums2[j] {
			nums[cnt] = nums1[i]
			i++
		} else {
			nums[cnt] = nums2[j]
			j++
		}
		cnt++
		if i == l1 {
			for j != l2 {
				nums[cnt] = nums2[j]
				j++
				cnt++
			}
		} else if j == l2 {
			for i != l1 {
				nums[cnt] = nums1[i]
				i++
				cnt++
			}
		}
	}
	return findMedian(nums)
}
func main() {
	// fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	// fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
}

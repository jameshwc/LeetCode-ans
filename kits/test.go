package main

import (
	"fmt"

	"github.com/jameshwc/LeetCode-ans/kits/data_structure/heap"
)

func main() {
	nums := []int{1, 7, 2, 3, 5, 10, 8, 9, 12, 3, 2, 4}
	h := heap.NewHeap(20, true)
	for i := range nums {
		h.Insert(nums[i])
	}
	for !h.IsEmpty() {
		fmt.Println(h.Pop())
	}
	// sorting.HeapSort(nums)
	// fmt.Println(nums)
}

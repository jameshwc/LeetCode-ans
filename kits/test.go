package main

import (
	"fmt"
	"kits/sorting/sorting"
)

func main() {
	nums := []int{1, 7, 2, 3, 5, 10, 8, 9, 12, 3, 2, 4}
	sorting.HeapSort(nums)
	fmt.Println(nums)
}

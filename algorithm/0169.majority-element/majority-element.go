package main

import "fmt"

func majorityElement(nums []int) int {
	var candidate, count int
	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	testA, testB := []int{3, 2, 3}, []int{2, 2, 1, 1, 1, 2, 2}
	if majorityElement(testA) == 3 && majorityElement(testB) == 2 {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}

package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left != right {
		lowerBound := height[left]
		if height[left] < height[right] {
			left++
		} else {
			lowerBound = height[right]
			right--
		}
		cur := lowerBound * (right - left + 1)
		if ans < cur {
			ans = cur
		}
	}
	return ans
}

func main() {
	if maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49 {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong Answer!")
	}
}

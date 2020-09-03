package main

import "fmt"

func dailyTemperatures(T []int) []int {
	stackT := make([]int, len(T))
	stackID := make([]int, len(T))
	ans := make([]int, len(T))
	top := 0
	for i := range T {
		for top > 0 && T[i] > stackT[top-1] {
			ans[stackID[top-1]] = i - stackID[top-1]
			top--
		}
		stackT[top] = T[i]
		stackID[top] = i
		top++
	}
	return ans
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T)) // [1, 1, 4, 2, 1, 1, 0, 0]
}

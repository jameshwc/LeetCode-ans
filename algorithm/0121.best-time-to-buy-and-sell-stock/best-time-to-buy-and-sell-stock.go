package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	maxProfit, min := 0, 1<<20
	for i := range prices {
		if min > prices[i] {
			min = prices[i]
		}
		if v := prices[i] - min; v > maxProfit {
			maxProfit = v
		}
	}
	return maxProfit
}

func main() {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 5 || maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		fmt.Println("Wrong Answer")
	} else {
		fmt.Println("Accepted")
	}
}

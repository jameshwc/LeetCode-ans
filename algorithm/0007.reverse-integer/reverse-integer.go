package main

import (
	"fmt"
)

func reverse(x int) int {
	const Limit = 2147483648
	var ans int
	for x != 0 {
		ans = (ans + (x % 10)) * 10
		x = x / 10
	}
	ans = ans / 10
	if ans > Limit-1 || ans < -Limit {
		return 0
	}
	return ans
}

func main() {
	x := -123
	fmt.Println(reverse(x))
}

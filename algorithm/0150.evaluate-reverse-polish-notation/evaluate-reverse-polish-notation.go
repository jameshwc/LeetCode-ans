package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	ans := 0
	stack := make([]int, 1<<10)
	top := 0
	for i := range tokens {
		switch tokens[i] {
		case "+", "-", "*", "/":
			n1, n2 := stack[top-2], stack[top-1]
			switch tokens[i] {
			case "+":
				ans = n1 + n2
			case "-":
				ans = n1 - n2
			case "*":
				ans = n1 * n2
			case "/":
				ans = n1 / n2
			}
			stack[top-2] = ans
			top--
		default:
			n, _ := strconv.Atoi(tokens[i])
			stack[top] = n
			top++
		}
	}
	return stack[top-1]
}

func main() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}

package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := 0
	code := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, r := range s {
		switch r {
		case '[', '{', '(':
			stack[top] = r
			top++
		case ']', '}', ')':
			if top < 1 || code[r] != stack[top-1] {
				return false
			}
			top--
		}
	}
	return top == 0
}

func main() {
	if !isValid("()[]{}") || isValid("(]") || isValid("([)]") || !isValid("{[]}") {
		fmt.Println("Wrong Answer!")
	} else {
		fmt.Println("Accepted!")
	}
}

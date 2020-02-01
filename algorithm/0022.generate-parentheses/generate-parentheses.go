package main

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

func recurGenParenthesis(L int, s string, ans *[]string) {
	if len(s) == L {
		if isValid(s) {
			*ans = append(*ans, s)
		}
	} else {
		recurGenParenthesis(L, s+"(", ans)
		recurGenParenthesis(L, s+")", ans)
	}
}
func generateParenthesis(n int) []string {
	var ans []string
	recurGenParenthesis(n*2, "", &ans)
	return ans
}

func main() {
	// fmt.Println(generateParenthesis(10))
	generateParenthesis(20)
}

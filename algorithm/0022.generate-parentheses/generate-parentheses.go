package main

func recurGenParenthesis(left int, right int, s string, L int, ans *[]string) {
	if left+right == 2*L {
		*ans = append(*ans, s)
		return
	}
	if left < L {
		recurGenParenthesis(left+1, right, s+"(", L, ans)
	}
	if left > right {
		recurGenParenthesis(left, right+1, s+")", L, ans)
	}
}
func generateParenthesis(n int) []string {
	var ans []string
	recurGenParenthesis(0, 0, "", n, &ans)
	return ans
}

func main() {
	// fmt.Println(generateParenthesis(10))
	generateParenthesis(3)
}

package main

func recurGenParenthesis(left int, right int, s string, L int, ans *[]string) {
	if right > left {
		return
	}
	if left+right != L {
		if left > right {
			recurGenParenthesis(left, right+1, s+")", L, ans)
		}
		recurGenParenthesis(left+1, right, s+"(", L, ans)
	} else if left == right {
		*ans = append(*ans, s)
	}

}
func generateParenthesis(n int) []string {
	var ans []string
	recurGenParenthesis(0, 0, "", 2*n, &ans)
	return ans
}

func main() {
	// fmt.Println(generateParenthesis(10))
	generateParenthesis(3)
}

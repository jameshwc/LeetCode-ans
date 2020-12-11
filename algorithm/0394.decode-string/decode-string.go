package main

import "fmt"

func decodeString(s string) string {
	if s == "" {
		return s
	}
	stackCh := make([]string, len(s))
	stackCnt := make([]int, len(s))
	sum, top := 0, 0
	for i := range s {
		if s[i] <= '9' && s[i] >= '0' {
			sum = sum*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			top++
			stackCnt[top] = sum
			stackCh[top] = ""
			sum = 0
		} else if s[i] == ']' {
			temp := ""
			for t := 0; t < stackCnt[top]; t++ {
				temp = temp + stackCh[top]
			}
			top--
			stackCh[top] = stackCh[top] + temp
		} else {
			stackCh[top] = stackCh[top] + string(s[i])
		}
	}
	return stackCh[top]
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))     // "aaabcbc"
	fmt.Println(decodeString("3[a2[c]]"))      // "accaccacc"
	fmt.Println(decodeString("2[abc]3[cd]ef")) // "abcabccdcdcdef"
	fmt.Println(decodeString("abc3[cd]xyz"))   // "abccdcdcdxyz"
}

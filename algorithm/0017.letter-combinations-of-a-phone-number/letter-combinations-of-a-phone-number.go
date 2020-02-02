package main

import "fmt"

func letterCombinations(digits string) []string {
	code := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return nil
	}
	ans := []string{""}
	for _, c := range digits {
		dc, _ := code[c]
		tmp := []string{}
		for _, cc := range ans {
			for _, ccc := range dc {
				tmp = append(tmp, cc+ccc)
			}
		}
		ans = tmp
	}
	return ans
}

func main() {
	fmt.Println(letterCombinations("23"))
}

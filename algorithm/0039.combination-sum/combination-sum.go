package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	n := len(candidates)
	var recf func(out []int, target int, id int)
	recf = func(out []int, target int, id int) {
		if target < 0 {
			return
		} else if target == 0 {
			tmp := make([]int, len(out))
			copy(tmp, out)
			ans = append(ans, tmp)
		} else {
			for i := id; i < n; i++ {
				out = append(out, candidates[i])
				recf(out, target-candidates[i], i)
				out = out[:len(out)-1]
			}
		}
	}
	var out []int
	recf(out, target, 0)
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

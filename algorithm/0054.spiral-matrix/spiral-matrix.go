package main

import "fmt"

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix[0])
	n := len(matrix)
	isVisited := make([][]bool, n)
	for i := range isVisited {
		isVisited[i] = make([]bool, m)
	}
	di := 0
	si, sj := 0, 0
	var ans []int
	for i := 0; i < m*n; i++ {
		isVisited[si][sj] = true
		ans = append(ans, matrix[si][sj])
		ei, ej := si+dir[di%4][0], sj+dir[di%4][1]
		if ei >= 0 && ej >= 0 && ei < n && ej < m && !isVisited[ei][ej] {
			si, sj = ei, ej
		} else {
			di++
			si += dir[di%4][0]
			sj += dir[di%4][1]
		}
	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2},
		{3, 4},
	}))
}

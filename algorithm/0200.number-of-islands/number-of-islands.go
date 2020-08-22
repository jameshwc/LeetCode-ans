package main

import "fmt"

func numIslands(grid [][]byte) int {
	island := 0
	isVisited := make([][]bool, len(grid))
	for i := range isVisited {
		isVisited[i] = make([]bool, len(grid[i]))
	}
	var dfs func(int, int)
	dfs = func(i, j int) { // index
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[i]) || isVisited[i][j] || grid[i][j] == '0' {
			return
		}
		isVisited[i][j] = true
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	for i := range grid {
		for j := range grid[i] {
			if isVisited[i][j] == false && grid[i][j] == '1' {
				island += 1
				dfs(i, j)
			}
		}
	}
	return island
}

/* isVisited can be removed by replacing grid[i][j] to '0' */

func numIslands(grid [][]byte) int {
	island := 0
	var dfs func(int, int)
	dfs = func(i, j int) { // index
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[i]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				island++
				dfs(i, j)
			}
		}
	}
	return island
}

func main() {
	testA := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	testB := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	if numIslands(testA) == 1 && numIslands(testB) == 3 {
		fmt.Println("Accepted!")
	} else {
		fmt.Println("Wrong answer!")
	}
}

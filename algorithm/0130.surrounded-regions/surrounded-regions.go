package main

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	isVisited := make([][]bool, len(board))

	for i := range board {
		isVisited[i] = make([]bool, len(board[i]))
	}

	isBorder := false
	var points [][]int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == len(board) || j == len(board[i]) || isVisited[i][j] || board[i][j] == 'X' {
			return
		}
		isVisited[i][j] = true
		points = append(points, []int{i, j})
		if i == 0 || j == 0 || i+1 == len(board) || j+1 == len(board[i]) {
			isBorder = true
		}
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' || isVisited[i][j] {
				continue
			}
			points = [][]int{}
			isBorder = false
			dfs(i, j)
			if !isBorder {
				for p := range points {
					board[points[p][0]][points[p][1]] = 'O'
				}
			}
		}
	}
}

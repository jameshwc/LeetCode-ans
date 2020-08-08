package main

func visit(isVisited [][]bool, board[][]byte, word string, posX, posY, wordID int) bool {
	if wordID == len(word) {
		return true
	}
	if posX < 0 || posY < 0 || posX == len(isVisited) || posY == len(isVisited[posX]) || board[posX][posY] != word[wordID] || isVisited[posX][posY] {
		return false
	}
	isVisited[posX][posY] = true
	if visit(isVisited, board, word, posX-1, posY, wordID+1) || visit(isVisited, board, word, posX+1, posY, wordID+1) ||
		visit(isVisited, board, word, posX, posY-1, wordID+1) || visit(isVisited, board, word, posX, posY+1, wordID+1) {
				return true
			}
	isVisited[posX][posY] = false
	return false
}
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			isVisited := make([][]bool, len(board))
			for i := range isVisited {
				isVisited[i] = make([]bool, len(board[i]))
			}
			if visit(isVisited, board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

// Closure version: save time and space in an efficient way

func exist(board [][]byte, word string) bool {
	var dfs func(int, int, int) bool
	dfs = func(posX, posY, wordID int) bool {
		if wordID == len(word) {
			return true
		}
		if posX < 0 || posY < 0 || posX == len(board) || posY == len(board[posX]) || board[posX][posY] == 0 || board[posX][posY] != word[wordID] {
			return false
		}
		tmp := board[posX][posY]
		board[posX][posY] = 0
		if dfs(posX-1, posY, wordID+1) || 
			dfs(posX+1, posY, wordID+1) ||
			dfs(posX, posY-1, wordID+1) ||
			dfs(posX, posY+1, wordID+1) {
				return true
			}
		board[posX][posY] = tmp
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
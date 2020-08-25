package main

import "fmt"

type Trie struct {
	nodes [26]*Trie
	word string
}

/** Initialize your data structure here. */
func newTrie() *Trie {
	return new(Trie)
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	current := this
	for _, ch := range word {
		index := ch - 'a'
		if current.nodes[index] == nil {
			current.nodes[index] = &Trie{}
		}
		current = current.nodes[index]
	}
	current.word = word
}

func findWords(board [][]byte, words []string) []string {
	trie := newTrie()
	for i := range words {
		trie.Insert(words[i])
	}
	var result []string
	var dfs func(int, int, *Trie, map[int]bool)
	dfs = func(i, j int, cur *Trie, isVisited map[int]bool) {
		if i < 0 || j < 0 || i == len(board) || j == len(board[i]) || isVisited[i*10000+j] {
			return
		}
		idx := board[i][j]-'a'
		if cur.nodes[idx] == nil {
			return
		}
		cur = cur.nodes[idx]
		if cur.word != "" {
			result = append(result, cur.word)
			cur.word = ""
		}
		isVisited[i*10000+j] = true
		dfs(i-1, j, cur, isVisited)
		dfs(i, j-1, cur, isVisited)
		dfs(i+1, j, cur, isVisited)
		dfs(i, j+1, cur, isVisited)
		delete(isVisited, i*10000+j)
	}
	for i := range board {
		for j := range board[i] {
			dfs(i, j, trie, make(map[int]bool))
		}
	}
	return result
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}

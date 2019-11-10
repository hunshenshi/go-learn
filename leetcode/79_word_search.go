package main

import "fmt"

func main() {
	board := [][]byte{{'A','B','C','E'},
		{'S','F','E','S'},
		{'A','D','E','E'}}
	word := "ABCESEEEFS"
	fmt.Printf("%c\n", board[0][2])
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 ||
		word == "" || len(word) == 0{
		return false
	}
	m, n := len(board), len(board[0])
	// go中这种表达式该怎么写
	//checked := [m * n]bool{}
	checked := make([]bool, m*n)
	//var checked [m*n]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matchStr(board, i, j, checked, word, 0) {
				return true
			}
		}
	}
	return false
}

func matchStr(board [][]byte, i, j int, checked []bool, word string, index int) bool {
	if index >= len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || checked[i*len(board[0])+j] {
		return false
	}
	// 回溯
	if board[i][j] == word[index] {
		checked[i*len(board[0]) + j] = true
		if matchStr(board, i, j+1, checked, word, index + 1) ||
			matchStr(board, i+1, j, checked, word, index + 1) ||
			matchStr(board, i, j-1, checked, word, index + 1) ||
			matchStr(board, i-1, j, checked, word, index + 1){
			return true
		}
		checked[i*len(board[0]) + j] = false
	}
	return false
}
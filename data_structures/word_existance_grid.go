package main

import "fmt"

var board [][] byte

func main() {
	board := [][] byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res := helper(i, j, "ABFD", 0, board)
			if res {
				break
			}
		}
	}
}

func helper(r int, c int, in string, pos int, board [][]byte) bool {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}
	if pos == len(in)-1 && board[r][c] == in[pos] {
		fmt.Print("found the word")
		fmt.Print("Row: ", r)
		fmt.Print("Col: ", c)
		return true
	}
	return helper(r+1, c, in, pos+1, board) ||
		helper(r, c+1, in, pos+1, board)
}

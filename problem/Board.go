package problem

import "fmt"

/**
board[0] = append(board[0], byte('X'))
board[0] = append(board[0], byte('X'))
board[0] = append(board[0], byte('X'))
board[0] = append(board[0], byte('X'))
board[1] = append(board[1], byte('X'))
board[1] = append(board[1], byte('O'))
board[1] = append(board[1], byte('O'))
board[1] = append(board[1], byte('X'))
board[2] = append(board[2], byte('X'))
board[2] = append(board[2], byte('X'))
board[2] = append(board[2], byte('O'))
board[2] = append(board[2], byte('X'))
board[3] = append(board[3], byte('X'))
board[3] = append(board[3], byte('O'))
board[3] = append(board[3], byte('X'))
board[3] = append(board[3], byte('X'))

board[0] = append(board[0], byte('O'))
board[0] = append(board[0], byte('O'))
board[0] = append(board[0], byte('O'))
board[1] = append(board[1], byte('O'))
board[1] = append(board[1], byte('O'))
board[1] = append(board[1], byte('O'))
board[2] = append(board[2], byte('O'))
board[2] = append(board[2], byte('O'))
board[2] = append(board[2], byte('O'))

board[0] = append(board[0], byte('O'))
board[0] = append(board[0], byte('X'))
board[0] = append(board[0], byte('X'))
board[0] = append(board[0], byte('O'))
board[0] = append(board[0], byte('X'))
board[1] = append(board[1], byte('X'))
board[1] = append(board[1], byte('O'))
board[1] = append(board[1], byte('O'))
board[1] = append(board[1], byte('X'))
board[1] = append(board[1], byte('O'))
board[2] = append(board[2], byte('X'))
board[2] = append(board[2], byte('O'))
board[2] = append(board[2], byte('X'))
board[2] = append(board[2], byte('O'))
board[2] = append(board[2], byte('X'))
board[3] = append(board[3], byte('O'))
board[3] = append(board[3], byte('X'))
board[3] = append(board[3], byte('O'))
board[3] = append(board[3], byte('O'))
board[3] = append(board[3], byte('O'))
board[4] = append(board[4], byte('X'))
board[4] = append(board[4], byte('X'))
board[4] = append(board[4], byte('O'))
board[4] = append(board[4], byte('X'))
board[4] = append(board[4], byte('O'))


*/
func TestProblem() {
	length := 4
	board := make([][]byte, length)
	board[0] = append(board[0], byte('X'))
	board[0] = append(board[0], byte('X'))
	board[0] = append(board[0], byte('X'))
	board[0] = append(board[0], byte('X'))
	board[1] = append(board[1], byte('X'))
	board[1] = append(board[1], byte('O'))
	board[1] = append(board[1], byte('O'))
	board[1] = append(board[1], byte('X'))
	board[2] = append(board[2], byte('X'))
	board[2] = append(board[2], byte('X'))
	board[2] = append(board[2], byte('O'))
	board[2] = append(board[2], byte('X'))
	board[3] = append(board[3], byte('X'))
	board[3] = append(board[3], byte('O'))
	board[3] = append(board[3], byte('X'))
	board[3] = append(board[3], byte('X'))
	solve(board)
	fmt.Println(board)
}

func solve(board [][]byte) {
	lengthA := len(board)
	for i, a := range board {
		lengthB := len(a)
		for j, b := range a {
			if (i == 0 || i == lengthA-1 || j == 0 || j == lengthB-1) && b == 'O' {
				findO(i, j, board, lengthA, lengthB)
			}
		}
	}
	fmt.Println(board)
	//将O变成X
	for i, a := range board {
		for j, b := range a {
			if b == 'O' {
				board[i][j] = 'X'
			} else if b == '-' {
				board[i][j] = 'O'
			}
		}
	}
	fmt.Println(board)
}
func findO(i int, j int, board [][]byte, lengthA int, lengthB int) {
	if i >= 0 && i < lengthA && j >= 0 && j < lengthB {
		if board[i][j] == 'O' {
			board[i][j] = '-'
			findO(i-1, j, board, lengthA, lengthB)
			findO(i+1, j, board, lengthA, lengthB)
			findO(i, j-1, board, lengthA, lengthB)
			findO(i, j+1, board, lengthA, lengthB)
		} else {
			return
		}
	} else {
		return
	}
}

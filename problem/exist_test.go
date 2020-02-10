package problem

func exist(board [][]byte, word string) bool {
	v := len(board)
	h := len(board[0])
	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			if word[0] == board[i][j] {
				pos := make([][]int, v)
				for n := 0; n < v; n++ {
					pos[n] = make([]int, h)
				}
				if recursiveExist(board, pos, 0, 0, word, 0) {
					return true
				}
			}
		}
	}
	return false
}
func recursiveExist(board [][]byte, pos [][]int, i int, j int, word string, current int) bool {
	v := len(board)
	h := len(board[0])
	right, left, down, up := false, false, false, false
	if board[i][j] == word[current] {
		pos[i][j] = 1
		if current == len(word)-1 {
			return true
		} else {
			if i+1 < v && pos[i+1][j] != 1 {
				pos[i][j] = 1
				right = recursiveExist(board, pos, i+1, j, word, current+1)
			}
			if i-1 >= 0 && pos[i-1][j] != 1 {
				left = recursiveExist(board, pos, i-1, j, word, current+1)
			}
			if j+1 < h && pos[i][j+1] != 1 {
				down = recursiveExist(board, pos, i, j+1, word, current+1)
			}
			if j-1 >= 0 && pos[i][j-1] != 1 {
				up = recursiveExist(board, pos, i, j-1, word, current+1)
			}
			return right || left || down || up
		}
	}
	pos[i][j] = 0
	return false
}

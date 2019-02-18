package problem

func solve(board [][]byte) {
	for i, a := range board {
		if i == 0 {
			continue
		}
		if i == len(a)-1 {
			continue
		}
		for j, b := range a {
			if j == 0 {
				continue
			}
			if j == len(a)-1 {
				continue
			}
			if b == 'X' {
				continue
			} else {
				//获取所有相连的O
				var queue [][]int
				findO(i, j, board, queue)
			}
		}
	}
}

func findO(i int, j int, board [][]byte, queue [][]int) {
	//向上找
	for i > 0 {
		i--
		if board[i][j] == 'O' {
			queue[i][j] = 1
			findO(i, j, board, queue)
			continue
		} else {
			break
		}
	}
	//向左找
	for j > 0 {
		j--
		if board[i][j] == 'O' {
			queue[i][j] = 1
			findO(i, j, board, queue)
			continue
		} else {
			break
		}
	}
	//向下找
	for i < len(board[i]) {
		i++
		if board[i][j] == 'O' {
			queue[i][j] = 1
			findO(i, j, board, queue)
			continue
		} else {
			break
		}
	}
	//向右找
	for j < len(board[i]) {
		j++
		if board[i][j] == 'O' {
			queue[i][j] = 1
			findO(i, j, board, queue)
			continue
		} else {
			break
		}
	}
}

package baseMind

import (
	"fmt"
	"os"
)

var resultQueen = make([]int, 8)

func TestQueen() {
	cal8Queens(0)
}
func cal8Queens(row int) {
	if row == 8 {
		printQueens(resultQueen)
		os.Exit(0)
	}
	for column := 0; column < 8; column++ {
		//位置是否符合
		if isOk(row, column) {
			resultQueen[row] = column
			row++
			//考察下一行
			cal8Queens(row)
		}
	}
}

func printQueens(queens []int) {
	fmt.Println(queens)
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			if queens[j] == i {
				fmt.Print("Q")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println("")
	}
}

func isOk(row int, column int) bool {
	left := column - 1
	right := column + 1
	for i := row - 1; i >= 0; i-- {
		if resultQueen[i] == column {
			return false
		}
		if left > 0 {
			if resultQueen[i] == left {
				return false
			}
		}
		if right < 8 {
			if resultQueen[i] == right {
				return false
			}
		}
		left--
		right++
	}
	return true
}

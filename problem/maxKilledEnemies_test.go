package problem

import "testing"

func maxKilledEnemies(grid [][]byte) int {
	maxEnemies := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if string(grid[i][j]) == "E" {
				left := j
				right := j
				for ; left >= 0; left-- {
					if string(grid[i][left]) == "W" {
						break
					}
					if string(grid[i][left]) != "W" && string(grid[i][left]) != "E" {
						if i > 0 {
							grid[i][left]++
							if string(grid[i-1][left]) != "W" && string(grid[i-1][left]) != "E" {
								grid[i][left] = grid[i][left] + grid[i-1][left]
							}
						} else {
							grid[i][left]++
						}
						if maxEnemies < int(grid[i][left]-48) {
							println(grid[i][left])
							maxEnemies = int(grid[i][left] - 48)
						}
					}
				}
				for ; right <= len(grid); right++ {
					if string(grid[i][right]) == "W" {
						break
					}
					if string(grid[i][right]) != "W" && string(grid[i][right]) != "E" {
						if i > 0 {
							grid[i][right]++
							if string(grid[i-1][right]) != "W" && string(grid[i-1][right]) != "E" {
								grid[i][right] = grid[i][right] + grid[i-1][right]
							}
						} else {
							grid[i][right]++
						}
						if maxEnemies < int(grid[i][right]-48) {
							println(grid[i][right])
							maxEnemies = int(grid[i][right] - 48)
						}
					}
				}
			}
		}
	}
	return maxEnemies
}
func TestMaxKilledEnemies(t *testing.T) {
	grid := [][]byte{[]byte("0E00"), []byte("E0WE"), []byte("0E00")}
	println(maxKilledEnemies(grid))
}

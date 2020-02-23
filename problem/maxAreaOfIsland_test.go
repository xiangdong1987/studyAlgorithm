package problem

import (
	"fmt"
	"testing"
)

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	if len(grid) < 1 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res := dfsIsland(grid, i, j)
				if max < res {
					max = res
				}
			}
		}
	}
	fmt.Println(grid)
	return max
}

func dfsIsland(grid [][]int, i int, j int) int {
	max := 1
	v := len(grid)
	h := len(grid[0])
	grid[i][j] = 0
	//向上
	if i-1 >= 0 && grid[i-1][j] == 1 {
		grid[i-1][j] = 0
		max += dfsIsland(grid, i-1, j)
	}
	//向下
	if i+1 < v && grid[i+1][j] == 1 {
		grid[i+1][j] = 0
		max += dfsIsland(grid, i+1, j)
	}
	//向左
	if j-1 >= 0 && grid[i][j-1] == 1 {
		grid[i][j-1] = 0
		max += dfsIsland(grid, i, j-1)
	}
	if j+1 < h && grid[i][j+1] == 1 {
		grid[i][j+1] = 0
		max += dfsIsland(grid, i, j+1)
	}
	return max
}

func TestMaxAreaOfIsland(t *testing.T) {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1}, {1, 1, 1},
	}))
}

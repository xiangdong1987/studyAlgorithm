package problem

func numIslands(grid [][]byte) int {
	h := len(grid)
	if h <= 0 {
		return 0
	}
	v := len(grid[0])
	sum := 0
	for i := 0; i < h; i++ {
		for j := 0; j < v; j++ {
			if grid[i][j] == byte('1') {
				sum++
				dfs(grid, i, j)
			}
		}
	}
	return sum
}

func dfs(grid [][]byte, i int, j int) {
	h := len(grid)
	v := len(grid[0])
	grid[i][j] = byte('0')
	if i-1 >= 0 && grid[i-1][j] == byte('1') {
		dfs(grid, i-1, j)
	}
	if i+1 < h && grid[i+1][j] == byte('1') {
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == byte('1') {
		dfs(grid, i, j-1)
	}
	if j+1 < v && grid[i][j+1] == byte('1') {
		dfs(grid, i, j+1)
	}
}

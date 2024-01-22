package datastruct

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	res := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				res++
				numIslandsDfs(grid, r, c)
			}
		}
	}
	return res
}
func numIslandsDfs(grid [][]byte, r, c int) {
	//超界
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	if grid[r][c] == '2' {
		return
	}
	grid[r][c] = '2'
	numIslandsDfs(grid, r-1, c)
	numIslandsDfs(grid, r+1, c)
	numIslandsDfs(grid, r, c-1)
	numIslandsDfs(grid, r, c+1)
}

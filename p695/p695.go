package p695

func Run(grid [][]int) int {
	return maxAreaOfIsland(grid)
}

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	if rows < 1 {
		return 0
	}
	cols := len(grid[0])
	if cols < 1 {
		return 0
	}
	max := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			tmp := dfs(i, j, grid)
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

func dfs(r, c int, grid [][]int) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	return 1 + dfs(r+1, c, grid) + dfs(r-1, c, grid) + dfs(r, c+1, grid) + dfs(r, c-1, grid)
}

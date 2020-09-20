package main

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	var count int

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == '1' {
				count++
				dfs(x, y, grid)
			}
		}
	}
	return count
}

func dfs(x, y int, grid [][]byte) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(x-1, y, grid)
	dfs(x+1, y, grid)
	dfs(x, y-1, grid)
	dfs(x, y+1, grid)
}

func numIslands(grid [][]byte) int {
	if len(grid) < 1 {
		return 0
	}
	var (
		n   int
		m   int
		ret = 0
	)
	n = len(grid)
	m = len(grid[0])

	var dfsMarking func(x, y int)
	dfsMarking = func(x, y int) {
		if x < n && x > 0 && y < m && y > 0 && grid[x][y] == '1' {
			grid[x][y] = '0'
			dfsMarking(x, y+1)
			dfsMarking(x+1, y)
			dfsMarking(x, y-1)
			dfsMarking(x-1, y)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				ret += 1
				dfsMarking(i, j)
			}
		}
	}
	return ret
}

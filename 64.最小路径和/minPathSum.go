package minPathSum

//依然是dp
func minPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[row-1][col-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
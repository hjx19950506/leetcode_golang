package solveSudoku

//回溯法
//特别注意二维数组的复制
var res [][]byte

func solveSudoku(board [][]byte)  {

	res = make([][]byte, 9)
	for i := 0; i < 9; i++ {
		res[i] = make([]byte, 9)
	}
	dfs(board, 0)
	for i := 0; i < 9; i++ {
		copy(board[i], res[i])
	}
}


func dfs(board [][]byte, k int) {
	if k == 81 {
		for i := 0; i < 9; i++ {
			copy(res[i], board[i])
		}
		return
	}
//	row := k / 9
//	col := k % 9
	if board[k/9][k%9] == '.' {
		for i := 1; i <= 9; i++ {
			board[k/9][k%9] = byte('0' + i)
			if !isNowValid(board, k) {
				board[k/9][k%9] = '.'
				continue
			}

			dfs(board, k+1)
			board[k/9][k%9] = byte('.')
		}
	} else {
		dfs(board, k+1)
	}
}
//比起36题做了简化
//只需检测当前行、列、块
func isNowValid(board [][]byte, k int) bool {
	Logs := make([]int, 3)

	row := k / 9
	col := k % 9
	block := row/3*3 + col/3
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' {
			x := board[row][i] - '1'
			if Logs[0]&(1<<x) != 0 {
				return false
			} else {
				Logs[0] |= 1 << x
			}
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' {
			x := board[i][col] - '1'
			if Logs[1]&(1<<x) != 0 {
				return false
			} else {
				Logs[1] |= 1 << x
			}
		}
	}
	row = block / 3 * 3
	col = block % 3 * 3
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if board[i][j] != '.' {
				x := board[i][j] - '1'
				if Logs[2]&(1<<x) != 0 {
					return false
				} else {
					Logs[2] |= 1 << x
				}
			}
		}
	}
	return true
}
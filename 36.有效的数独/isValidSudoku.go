package isValidSudoku


//使用三个数组分别记录行、列、块
//扫描全表
//若当前数字已记录 则无效
//否则记录当前数字
/*func isValidSudoku(board [][]byte) bool {
	rowLog := make([]int, 9)
	colLog := make([]int, 9)
	blockLog := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				x := board[i][j] - '1'
				//行判断
				if (rowLog[i] & (1 << x)) != 0 {
					return false
				} else {
					rowLog[i] = rowLog[i] | (1 << x)
				}
				//列判断
				if (colLog[j] & (1 << x)) != 0 {
					return false
				} else {
					colLog[j] = colLog[j] | (1 << x)
				}
				//块判断
				k := i / 3 * 3 + j / 3
				if (blockLog[k] & (1 << x)) != 0 {
					return false
				} else {
					blockLog[k] = blockLog[k] | (1 << x)
				}
			}
		}
	}
	return true
}
*/

//优化版本，简化了流程
//思路和上面的一样
func isValidSudoku(board [][]byte) bool {
	Logs := make([][]int, 3)
	for i := 0; i < 3; i++ {
		Logs[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				x := board[i][j] - '1'
				k := i / 3 * 3 + j / 3
				r := []int{i, j, k}
				for index, val := range r {
					if (Logs[index][val] & (1 << x)) != 0 {
						return false
					} else {
						Logs[index][val] = Logs[index][val] | (1 << x)
					}
				}
			}
		}
	}
	return true
}
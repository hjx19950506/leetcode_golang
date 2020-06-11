package exist

//全程使用log二维布尔数组记录当前每个位置的字母是否已被使用
//遍历board找到匹配word第一个字母的元素位置
//然后进入dfs
var rowCount int
var colCount int
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	rowCount = len(board)
	if rowCount == 0 {
		return false
	}
	colCount = len(board[0])
	log := make([][]bool, rowCount)
	for i := range log {
		log[i] = make([]bool, colCount)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				log[i][j] = true
				if dfs(board, word, log, i, j, 1) {
					return true
				}
				log[i][j] = false
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, log [][]bool, row, col, k int) bool {
	if k == len(word) {
		return true
	}
	res := false

	dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	for i := 0; i < 4; i++ {
		rowNew, colNew := row+dir[i][0], col+dir[i][1]
		if !(rowNew >= 0 && rowNew < rowCount && colNew >= 0 && colNew < colCount) {
			continue
		}
		if board[rowNew][colNew] == word[k] && log[rowNew][colNew] == false {
			log[rowNew][colNew] = true
			res = res || dfs(board, word, log, rowNew, colNew, k+1)
			log[rowNew][colNew] = false
		}
	}
	return res
}
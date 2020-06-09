package searchMatrix

//从左下角开始找
//当前数字比target大则往上移动，小则往右移动
func searchMatrix(matrix [][]int, target int) bool {
	rowNum := len(matrix)
	if rowNum == 0 {
		return false
	}
	colNum := len(matrix[0])


	for i, j := rowNum - 1, 0; i >= 0 && j < colNum; {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
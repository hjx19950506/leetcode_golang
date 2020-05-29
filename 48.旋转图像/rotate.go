package rotate

//先转置，然后水平翻转
func rotate(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}
	rowLen := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < rowLen; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < rowLen/2; j++ {
			matrix[i][j], matrix[i][rowLen-j-1] = matrix[i][rowLen-j-1], matrix[i][j]
		}
	}
}
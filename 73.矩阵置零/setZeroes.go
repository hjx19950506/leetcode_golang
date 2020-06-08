package setZeroes

//一次遍历，使用首行与首列记录当前行列是否需要置零
//再遍历一次完成置零，要注意首行与首列最后处理
func setZeroes(matrix [][]int)  {
	rowCount := len(matrix)
	if rowCount == 0 {
		return
	}
	colCount := len(matrix[0])
	if colCount == 0 {
		return
	}
	firstRow, firstCol := 1, 1

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = 0
				}
				if j == 0 {
					firstCol = 0
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rowCount; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < colCount; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < colCount; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < rowCount; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if firstCol == 0 {
		for i := 0; i < rowCount; i++ {
			matrix[i][0] = 0
		}
	}
	if firstRow == 0 {
		for j := 0; j < colCount; j++ {
			matrix[0][j] = 0
		}
	}
	return
}
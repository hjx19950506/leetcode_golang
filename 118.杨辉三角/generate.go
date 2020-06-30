package generate

//按层遍历即可
//每一层首位尾元素为1
//其他元素为上一行相邻两个元素之和
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}
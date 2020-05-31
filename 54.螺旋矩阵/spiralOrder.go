package spiralOrder


//直观的进行模拟
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	l, r := 0, len(matrix[0]) - 1
	u, d := 0, len(matrix) -1
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[u][i])
		}
		if u++; u > d {
			break
		}
		for i := u; i <= d; i++ {
			res = append(res, matrix[i][r])
		}
		if r--; r < l {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, matrix[d][i])
		}
		if d--; d < u {
			break
		}
		for i := d; i >= u; i-- {
			res = append(res, matrix[i][l])
		}
		if l++; l > r {
			break
		}
	}
	return res
}
package solveNQueens

import (
	"bytes"
)

//dfs问题
func solveNQueens(n int) [][]string {
	res := [][]string{}
	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = bytes.Repeat([]byte("."), n)
	}

	dfs(n, 0, &res, buf)
	return res
}

func dfs(n, k int, res *[][]string, buf [][]byte){
	if k == n {
		*res = append(*res, dump(buf))
		return
	}
	for i := 0; i < n; i++ {
		if isValid(buf, k, i) {
			buf[k][i] = 'Q'
			dfs(n, k+1, res, buf)
			buf[k][i] = '.'
		}
	}
}

func isValid(buf [][]byte, row, col int) bool {
	for i := 0; i < row; i++ {
		for j := 0; j < len(buf); j++ {
			if buf[i][j] == 'Q' && (j == col || i-row == j-col || i-row == col-j) {
				return false
			}
		}
	}
	return true
}

func dump(src [][]byte) []string {
	dst := make([]string, len(src))
	for i, v := range src {
		dst[i] = string(v)
	}
	return dst
}
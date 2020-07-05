package minimumTotal

//自下而上做dp
func minimumTotal(triangle [][]int) int {
	rowNum := len(triangle)
	dp := make([]int, rowNum+1)
	for i := rowNum - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/* 由上而下
//dp
//使用一个pre数组存储到达上一行为止，
//每个结点所需要的路径和
//结合当前行的每个结点，
//得到cur数组，并更新pre
func minimumTotal(triangle [][]int) int {
	rowNum := len(triangle)
	if rowNum == 0 {
		return 0
	}
	pre, cur := make([]int, rowNum), make([]int, rowNum)
	pre[0] = triangle[0][0]
	for i := 1; i < rowNum; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				cur[j] = pre[j]
			} else if j == len(triangle[i]) - 1 {
				cur[j] = pre[j-1]
			} else {
				cur[j] = min(pre[j], pre[j-1])
			}
			cur[j] += triangle[i][j]
		}
		copy(pre, cur)
	}
	res := cur[0]
	for i := range cur {
		if cur[i] < res {
			res = cur[i]
		}
	}
	return res
}
*/
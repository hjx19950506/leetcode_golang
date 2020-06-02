package generateMatrix

//类似54题
//先把空间创建好
//模拟按层循环去填写
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}
	k := 1
	l, r, u, d := 0, n-1, 0, n-1

	for {
		//上
		for i := l; i <= r; i++ {
			res[u][i] = k
			k++
		}
		if u++; u > d {
			break
		}
		//右
		for i := u; i <= d; i++ {
			res[i][r] = k
			k++
		}
		if r--; r < l {
			break
		}
		//下
		for i := r; i >= l; i-- {
			res[d][i] = k
			k++
		}
		if d--; d < u {
			break
		}
		//左
		for i := d; i >= u; i-- {
			res[i][l] = k
			k++
		}
		if l++; l > r {
			break
		}
	}
	return res
}
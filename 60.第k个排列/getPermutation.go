package getPermutation

import "strconv"

//康托逆展开
func getPermutation(n int, k int) string {
	res := ""
	log, div := make([]int, n), make([]int, n)
	//计算阶乘结果
	div[n-1] = 1
	for i := n-2; i >= 0; i-- {
		div[i] = div[i+1] * (n-1-i)
	}
	k--	//康托排序从0开始
	for i := 0; i < n; i++ {
		curRank := k / div[i] + 1
		k %= div[i]
		for j, count := 0, 0; j < n; j++ {
			if log[j] == 0 {
				count++
			}
			if count == curRank {
				res += strconv.Itoa(j+1)
				log[j] = 1
				break
			}
		}
	}
	return res
}


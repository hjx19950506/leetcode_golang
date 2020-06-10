package combine

//从buf的最右开始，递增该位数字得到buf存入res
//当当前位数字达到最大时，向左移一位重复以上步骤
func combine(n int, k int) [][]int {
	res := [][]int{}
	if k == 0 {
		return res
	}
	buf := make([]int, k)

	for i := 0; i >= 0; {
		buf[i]++
		if buf[i] > n {
			i--
		} else if i == k - 1 {
			res = append(res, dump(buf[:k]))
		} else {
			i++
			buf[i] = buf[i-1]
		}
	}

	return res
}

/*
func combine(n int, k int) [][]int {
	res := [][]int{}
	buf := []int{}
	for i := 1; i <= k; i++ {
		buf = append(buf, i)
	}
	buf = append(buf, n+1)
	for j := 0; j < k; {
		res = append(res, dump(buf[:k]))
		j = 0
		for j < k && buf[j+1] == buf[j] + 1 {
			buf[j] = j + 1
			j++
		}
		buf[j] = buf[j] + 1
	}
	return res
}*/


/*
func combine(n int, k int) [][]int {

	log := make([]int, n)
	res := make([][]int, 0)
	buf := make([]int, 0)

	dfs(log, buf, &res, n, k)
	return res
}

func dfs(log, buf []int, res *[][]int, n, k int)  {
	if k == 0 {
		*res = append(*res, dump(buf))
		return
	}

	i := 0
	if len(buf) == 0 {
		i = 0
	} else {
		i = buf[len(buf)-1]
	}

	for ; i < n; i++ {
		if log[i] == 0 {
			log[i] = 1
			buf = append(buf, i+1)
			dfs(log, buf, res, n, k-1)
			buf = buf[0:len(buf)-1]
			log[i] = 0
		}
	}
}
*/
func dump(buf []int) []int {
	dst := make([]int, len(buf))
	copy(dst, buf)
	return dst
}

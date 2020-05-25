package combinationSum


//dfs
//要注意slice是指针传递
//所以存入result的时候要复制一份
//copy复制的元素个数为两个参数个数少的那个
var result [][]int
func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	curBuf := make([]int, 0)
	dfs(candidates, target, curBuf, 0)
	return result
}

func dfs(candidates []int, target int, curBuf []int, k int) {
	if target == 0 {
		temp := make([]int, len(curBuf))
		copy(temp, curBuf)
		result = append(result, temp)
		return
	}
	if target < 0 {
		return
	}

	for i := k; i < len(candidates); i++ {
		if target < candidates[i] {
			continue
		}
		curBuf = append(curBuf, candidates[i])
		dfs(candidates, target - candidates[i], curBuf, i)
		curBuf = curBuf[:len(curBuf)-1]
	}
}



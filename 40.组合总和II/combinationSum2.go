package combinationSum2

import "sort"

//与39类似
//为了剪枝提前排序
//遍历时当前元素与前一元素相同则继续往前，因为必被包含了
//此外，在继续深入使k+1以避免使用同一元素
var result [][]int
func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	curBuf := make([]int, 0)
	sort.Ints(candidates)
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
		if i > k && candidates[i] == candidates[i-1] {
			continue
		}
		//有序，直接可以返回
		if target < candidates[i] {
			return
		}
		curBuf = append(curBuf, candidates[i])
		dfs(candidates, target - candidates[i], curBuf, i+1)
		curBuf = curBuf[:len(curBuf)-1]
	}
}



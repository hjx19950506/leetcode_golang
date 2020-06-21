package subsetsWithDup

import "sort"

//与78题类似,但需要去重
//在内层循环使用j != 0 && nums[j] == nums[j-1] && (1<<(j-1) & i == 0)判断
//符合该条件时，说明前一个数与当前数相同，且前一个数不在子集中，
//那么当前方法必被遍历过了
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	numsLen := len(nums)
	maxMask := 1 << numsLen - 1

	for i := 0; i <= maxMask; i++ {
		buf := []int{}
		for j := 0; j < numsLen; j++ {
			if 1<<j & i != 0 {
				if j != 0 && nums[j] == nums[j-1] && (1<<(j-1) & i == 0) {
					goto Loop
				}
				buf = append(buf, nums[j])
			}
		}
		res = append(res, buf)
		Loop:
	}
	return res
}
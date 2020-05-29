package permuteUnique


//在46题基础上剪枝
//使用一个map记录当前位置上是否已经使用过当前遍历的元素
//如是，则是子集，一定是被遍历过了
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, &res,0)
	return res
}

func dfs(nums []int, res *[][]int, k int) {
	if k == len(nums) {
		*res = append(*res, dump(nums))
	}
	m := map[int]bool{}
	for i := k; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		//该句不能放在两次交换之间
		m[nums[i]] = true
		nums[i], nums[k] = nums[k], nums[i]
		dfs(nums, res, k+1)
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func dump(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

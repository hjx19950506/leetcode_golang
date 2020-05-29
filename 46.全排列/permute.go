package permute

//改进的dfs
//使用O(1)额外空间
//不借助额外数组暂存排列
//而是在原数组上做交换
func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, &res, 0)
	return res
}

func dfs(nums []int, res *[][]int, k int) {
	if k == len(nums) {
		*res = append(*res, dump(nums))
	}
	for i := k; i < len(nums); i++ {
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

/*
//基础的dfs题
var numsLen int
var results [][]int

func permute(nums []int) [][]int {
	numsLen = len(nums)
	results = make([][]int, 0)
	buff := make([]int, 0)
	log := make([]bool, numsLen)
	dfs(nums, buff, log, 0)
	return results
}

func dfs(nums, buff []int, log []bool, k int) {

	if k == numsLen {
		temp := make([]int, numsLen)
		copy(temp, buff)
		results = append(results, temp)
	}
	for i := 0; i < numsLen; i++ {
		if !log[i] {
			buff = append(buff, nums[i])
			log[i] = true
			dfs(nums, buff, log, k+1)
			buff = buff[:len(buff)-1]
			log[i] = false
		}
	}
}
*/

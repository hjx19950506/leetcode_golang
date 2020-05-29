package permute


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
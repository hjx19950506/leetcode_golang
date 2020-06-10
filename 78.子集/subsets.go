package subsets

//使用掩码的方式找出所有子集
//掩码从0...0递增到1...1
//分别代表所有数字都不要和都要
func subsets(nums []int) [][]int {
	res := [][]int{}
	numsLen := len(nums)
	maxMask := 1 << numsLen - 1

	for i := 0; i <= maxMask; i++ {
		buf := []int{}
		for j := 0; j < numsLen; j++ {
			if 1<<j & i != 0 {
				buf = append(buf, nums[j])
			}
		}
		res = append(res, buf)
	}
	return res
}
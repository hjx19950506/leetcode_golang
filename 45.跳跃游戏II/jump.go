package jump

//向前多看一步
//尽量走远
//走的近的方式必是走的远的方式的子集
func jump(nums []int) int {
	res := 0
	for i := 0; i < len(nums) - 1; {
		max := 0
		index := i
		for j := 1; i+j < len(nums) && j <= nums[i]; j++ {
			//提前结束
			if i + j >= len(nums) - 1 {
				return res+1
			}
			if j + nums[i+j] > max {
				max = j + nums[i+j]
				index = j
			}
		}
		i += index
		res++
	}
	return res
}
package canJump

func canJump(nums []int) bool {
	for i := 0; i < len(nums) - 1; {
		max := 0
		index := i
		for j := 1; i+j < len(nums) && j <= nums[i]; j++ {
			//提前结束
			if i + j >= len(nums) - 1 {
				return true
			}
			if j + nums[i+j] > max {
				max = j + nums[i+j]
				index = j
			}
		}
		i += index
		if nums[i] == 0 {
			return false
		}
	}
	return true
}
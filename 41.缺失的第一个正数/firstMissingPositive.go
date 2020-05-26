package firstMissingPositive

//一次遍历将元素尽量归位
//第二次查找答案
func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for nums[i] != i + 1 && nums[i] > 0 && nums[i] <= numsLen && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}

	}
	for i := 0; i < numsLen; i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}
	return numsLen+1
}
package removeDuplicates

//一次遍历，双指针
//遍历时是阶段式的
//每次使用r指针记录一段相同数字的长度
//然后在l指针处打印最多2个r指针处的数字
func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen < 3 {
		return numsLen
	}
	l, r := 0, 0
	for r < numsLen {
		count := 1
		for r < numsLen - 1 && nums[r] == nums[r+1] {
			count++
			r++
		}
		if count > 2 {
			count = 2
		}
		for i := 0; i < count; i++ {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	nums = nums[:l]
	return len(nums)
}
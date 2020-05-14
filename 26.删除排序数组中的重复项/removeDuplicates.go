package removeDuplicates

//双指针
//一个指针指向当前不重复数组末尾
//另一个去寻找新的不重复数字
//并将其覆盖到第一个指针指向的下一个数
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	i := 1
	for j := 1; j < length; j++ {
		for nums[j] == nums[i-1] {
			j++
			if j > length - 1 {
				return i
			}
		}
		nums[i] = nums[j]
		i++
	}
	return i
}
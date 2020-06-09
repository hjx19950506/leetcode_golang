package sortColors

//左右个一个指针分别指向最后一个0和第一个2
//当前指针去遍历数组
//为0则与左指针对调，为2则与右指针对调
func sortColors(nums []int)  {
	numsLen := len(nums)

	for left, right, cur := 0, numsLen - 1, 0; cur <= right; {
		if nums[cur] == 0 {
			nums[left], nums[cur] = nums[cur], nums[left]
			left++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		} else {
			cur++
		}
	}
	return
}
package nextPermutation


//从右往左找到第一个比它右边元素小的元素
//交换该元素和它右边元素中比它大的最小数
//这样该元素右边仍是降序
//将它右边逆排即可
func nextPermutation(nums []int)  {
	numsLen := len(nums)
	if numsLen == 0 {
		return
	}
	i := 0
	for i = numsLen - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j := numsLen - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums) - 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

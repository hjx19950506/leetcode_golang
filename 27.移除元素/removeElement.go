package removeElement

//和26题思路类似
//前后指针
func removeElement(nums []int, val int) int {
	length := len(nums)
	i, j := 0, 0
	for i < length || j < length {
		for j < length && nums[j] == val {
			j++
		}
		if j >= length {
			break;
		}
		nums[i] = nums[j]
		i++
		j++
	}
	return i
}
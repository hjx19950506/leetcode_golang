package searchInsert

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 || target < nums[0] {
		return 0
	}
	if target > nums[numsLen-1] {
		return numsLen
	}
	l, r := 0, numsLen
	mid := 0
	for l < r {
		mid = (l + r) / 2
		if nums[mid] == target || (nums[mid] > target && nums[mid-1] < target) {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return mid
}
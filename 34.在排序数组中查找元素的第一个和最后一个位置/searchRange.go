package searchRange

//编写两个函数
//分别寻找第一个和最后一个target的索引
//方式都是二分查找
func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	l, r := 0, numsLen
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}

func findLast(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	l, r := 0, numsLen
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target && (mid == numsLen - 1 || nums[mid+1] > target) {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
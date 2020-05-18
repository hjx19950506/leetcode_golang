package search

//二分查找+递归
//如果左边有序，且target在左边最大最小值间，则继续在左边二分查找，否则递归右边
//如果右边有序，且target在右边最大最小值间，则继续在右边二分查找，否则递归左边
func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	if numsLen == 1 && nums[0] != target {
		return -1
	}
	l, r := 0, numsLen
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[l] < nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				r = mid
			} else {
				l = mid + 1
				break
			}
		} else {
			if target > nums[mid] && target <= nums[r-1] {
				l = mid + 1
			} else {
				r = mid
				break
			}
		}
	}

	insideIndex := search(nums[l:r], target)
	if insideIndex == -1 {
		return -1
	}
	return l + insideIndex
}

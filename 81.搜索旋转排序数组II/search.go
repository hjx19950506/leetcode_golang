package search

//与33题思路一致
//二分查找+递归
//如果左边有序，且target在左边最大最小值间，则继续在左边二分查找，否则递归右边
//如果右边有序，且target在右边最大最小值间，则继续在右边二分查找，否则递归左边
//由于nums有重复，所以每次保证左半部分没有重复
func search(nums []int, target int) bool {
	numsLen := len(nums)
	if numsLen == 0 {
		return false
	}
	if numsLen == 1 {
		return nums[0] == target
	}
	l, r := 0, numsLen
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] {
			l++
			continue
		}
		if nums[l] < nums[mid] {
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

	return search(nums[l:r], target)
}
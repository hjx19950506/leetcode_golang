package threeSumClosest

import "sort"

//思路与15题相似
//只需在每次移动双指针对的时候观察一下三数和就行
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)

	for i := 0; i < length-2; i++ {
		for l, r := i+1, length-1; l < r; {
			temp := nums[i] + nums[l] +nums[r]
			if abs(temp - target) < abs(res - target) {
				res = temp
			}
			if temp < target {
				l++
			} else if temp > target {
				r--
			} else {
				return temp
			}
			for l > i + 1 && l < r && nums[l] == nums[l-1] {
				l++
			}
			for r > l && r < length - 1 && nums[r] == nums[r+1] {
				r--
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
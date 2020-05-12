package fourSum

import "sort"

//思路与15三数之和类似
//再套一层循环即可
//注意去重 每次指针变换都要判断
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < length - 3; i++ {
		for i > 0 && i < length - 3 && nums[i] == nums[i-1] {
			i++
		}
		for j := i + 1; j < length - 2; j++ {
			for j > i + 1 && j < length - 2 && nums[j] == nums[j-1] {
				j++
			}
			for l, r := j + 1, length-1; l < r; {
				tempSum := nums[i] + nums[j] + nums[l] + nums[r]
				if tempSum < target {
					l++
				} else if tempSum > target {
					r--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
				for l > j + 1 && l < r && nums[l] == nums[l-1] {
					l++
				}
				for r > l && r < length - 1 && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}

	return res
}
package threeSum

import "sort"

//先排序 O(nlogn)
//然后一个指针遍历数组O(n)
//再用一对双指针遍历O(n)
//总时间复杂度为O(n*n)
func threeSum(nums []int) [][]int {
	res := [][]int{}
	length := len(nums)
	sort.Ints(nums)

	for i := 0; i < length - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		for l, r := i + 1, length - 1; l < r; {
			tmpSum := nums[i] + nums[l] + nums[r]
			if tmpSum < 0 {
				//总和小了左指针继续往左
				l++
			} else if tmpSum > 0 {
				//综合大了右指针继续往右
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
			for l < length && l > i + 1 && nums[l] == nums[l-1] {
				l++
			}
			for r < length - 1 && r > 1 && nums[r] == nums[r+1] {
				r--
			}
		}
	}

	return res
}
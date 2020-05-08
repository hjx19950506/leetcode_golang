package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//确保nums1元素不多于nums2
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	//二分查找
	iMin, iMax := 0, len(nums1)
	i := (iMin + iMax) / 2
	j := (len(nums1) + len(nums2) + 1) / 2 - i
	//for nums1[i-1] > nums2[j] || nums2[j-1] > nums1[i] {
	for iMin <= iMax {
		i = (iMin + iMax) / 2
		j = (len(nums1) + len(nums2) + 1) / 2 - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else {
			break
		}
	}

	//处理边界情况
	var leftMax, rightMin float64
	if i == 0 {
		leftMax = float64(nums2[j-1])
	} else if j == 0 {
		leftMax = float64(nums1[i-1])
	} else {
		leftMax = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
	}
	//总共奇数个数则返回
	if (len(nums1) + len(nums2)) % 2 == 1 {
		return leftMax
	}
	if i == len(nums1) {
		rightMin = float64(nums2[j])
	} else if j == len(nums2) {
		rightMin = float64(nums1[i])
	} else {
		rightMin = math.Min(float64(nums1[i]), float64(nums2[j]))
	}

	//leftMax,rightMin的平均值为中位数
	return (leftMax + rightMin) / 2
}

func main() {
	nums1 := []int{1,3,9,10}
	nums2 := []int{1,2,3,4,5,6,7,8,9,10}

	res := findMedianSortedArrays(nums2, nums1)
	print(res)
}
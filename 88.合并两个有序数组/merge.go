package merge

//两个指针分别指向两个数组的最后一个有效数字
//将较大的填到nums1的末尾
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i, j, k := m - 1, n - 1, m + n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 {
			if nums1[i] < nums2[j] {
				nums1[k] = nums2[j]
				j--
			} else {
				nums1[k] = nums1[i]
				i--
			}
		} else if i < 0 && j >= 0 {
			nums1[k] = nums2[j]
			j--
		} else if i >= 0 && j < 0 {
			nums1[k] = nums1[i]
			i--
		}
	}
}
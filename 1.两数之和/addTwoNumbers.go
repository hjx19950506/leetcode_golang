package addTwoNumbers

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i1, v := range nums {
		if i2, ok := m[target - v]; ok {
			return []int{i1, i2}
		}
		m[v] = i1
	}

	return nil
}
package romanToInt

func romanToInt(s string) int {
	res := 0
	m := make(map[string]int)
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	strings := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i, v := range strings {
		m[v] = nums[i]
	}

	i := 0
	for ; i < len(s) - 1; i++ {
		if v, ok := m[string(s[i]) + string(s[i+1])]; ok{
			//俩字母
			res += v
			i++
		} else {
			//一字母
			res += m[string(s[i])]
		}
	}
	if i == len(s) - 1 {
		res += m[string(s[i])]
	}

	return res
}
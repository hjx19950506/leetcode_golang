package intToRoman

/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

//依次减去尽量多的1000 900 500 400...
func intToRoman(num int) string {
	res := ""
	m := make(map[int]string)
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	strings := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i, v := range nums {
		m[v] = strings[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for num >= nums[i] {
			res += m[nums[i]]
			num -= nums[i]
		}
	}

	return res
}
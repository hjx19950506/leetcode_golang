package plusOne

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	res := make([]int, 0)
	carry := 1
	for i := digitsLen-1; i >= 0; i-- {
		k := digits[i] + carry
		digits[i] = k % 10
		carry = k / 10
		if carry == 0 {
			break
		}
	}
	if carry != 0 {
		res = append(res, carry)
	}
	res = append(res, digits...)
	return res
}
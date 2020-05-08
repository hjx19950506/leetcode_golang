package myAtoi

import "math"

//类似第7题的32位限制
//此外若想从for循环内的switch内直接跳出循环需要break+tag
func myAtoi(str string) int {
	res := int32(0)
	sign := int32(0)

	Loop:
	for _, v := range str {
		switch {
		case v == ' ':
			if sign != 0 {
				break Loop
			}
			continue
		case v == '-':
			if sign == 0 {
				sign = -1
			} else {
				break Loop
			}
		case v == '+':
			if sign == 0 {
				sign = 1
			} else {
				break Loop
			}
		case v >= '0' && v <= '9':
			if sign == 0 {
				sign = 1
			}
			tmp := (v - '0') * sign
			if res > math.MaxInt32 / 10 || (res == math.MaxInt32 / 10 && tmp > 7) {
				return math.MaxInt32
			}
			if res < math.MinInt32 / 10 || (res == math.MinInt32 / 10 && tmp < -8) {
				return math.MinInt32
			}
			res = res * 10 + tmp
		default:
			break Loop
		}
	}

	return int(res)
}

package divide

import "math"

//先将divisor不断翻倍
//找到使得divisor*2^n不小于dividend地最小n
//然后逐渐减小n看dividend-divisor*2^n是否大于0
//相应在res中加n
func divide(dividend int, divisor int) int {
	n1, n2 := int32(dividend), int32(divisor)
	if n1 == math.MinInt32 && n2 == -1 {
		return int(math.MaxInt32)
	}
	if n1 == math.MinInt32 && n2 == math.MinInt32 {
		return 1
	}
	res := int32(0)
	//对符号做处理
	if n1 == math.MinInt32 {
		if n2 < 0 {
			n1 -= n2
		} else {
			n1 += n2
		}
		res++
	}
	flag := (n1 > 0 && n2 > 0) || (n1 < 0 && n2 < 0)
	if n1 < 0 {
		n1 = -n1
	}
	if n2 < 0 {
		n2 = -n2
	}
	//求到最小的n
	n := int32(0)
	for int32(n2 << n) > 0 && n2 << n <= n1 {
		n++
	}
	n--
	//累加res
	for n >= 0 {
		if n2 << n <= n1 {
			n1 -= n2 << n
			res += 1 << n
		}
		n--
	}

	if !flag {
		return int(-res)
	}

	return int(res)
}
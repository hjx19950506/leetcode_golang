package reverse

import "math"

//题目要求只可用到int32
//golang的int类型不固定
//32/64位机上为4/8字节
//所以需要先做转换
func reverse(x int) int {
	y := int32(0)
	x32 := int32(x)

	for x32 != 0 {
		remainder := x32 % 10
		x32 /= 10
		//判断是否即将溢出
		if y > math.MaxInt32 / 10 || (y == math.MaxInt32 / 10 && remainder > 7) {
			return 0
		}
		if y < math.MinInt32 / 10 || (y == math.MinInt32 / 10 && remainder < -8) {
			return 0
		}

		y = y * 10 + remainder
	}

	return int(y)
}

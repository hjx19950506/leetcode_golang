package multiply

import "strconv"

//竖式乘法
//遍历num1和num2的每一位
//得到的结果个位十位存入resInt[i+j+1]、resInt[i+j]
//再遍历一次resInt转化为字符
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	resInt := make([]int, l1+l2)
	res := ""

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			temp := int((num1[i]-'0')*(num2[j]-'0'))
			unit := temp % 10
			ten := temp / 10
			resInt[i+j+1] += unit
			resInt[i+j] += ten
		}
	}
	for i := l1+l2-1; i > 0; i-- {

		resInt[i-1] += resInt[i] / 10
		resInt[i] %= 10
		res = strconv.Itoa(resInt[i]) + res
	}
	if resInt[0] != 0 {
		res = strconv.Itoa(resInt[0]) + res
	}
	return res
}
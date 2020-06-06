package addBinary

import "strings"

//使用补全较短字符串的高位0后
//从低到高计算结果和进位
func addBinary(a string, b string) string {
	res := ""
	aLen, bLen := len(a), len(b)
	if aLen < bLen {
		aLen, bLen = bLen, aLen
		a, b = b, a
	}
	resTemp := make([]byte, aLen)
	b = strings.Repeat("0", aLen - bLen) + b
	carry := uint8(0)
	for i := aLen-1; i >= 0; i-- {
		k := (a[i]-'0') + (b[i]-'0') + carry
		resTemp[i] = k % 2 + '0'
		carry = k / 2
	}
	if carry != 0 {
		res = "1"
	}
	res = res + string(resTemp)
	return res
}
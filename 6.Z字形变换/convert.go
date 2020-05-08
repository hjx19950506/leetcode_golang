package convert

import "math"

//计算变成z形后字符串行数，遍历s将字母放入对应行，使用一个uping标记方向，
//最后将所有字符串拼接
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	rowNum := int(math.Min(float64(length), float64(numRows)))
	strings := make([][]byte, rowNum)
	nowRow := 0
	uping := false

	for _, v := range s {
		strings[nowRow] = append(strings[nowRow], byte(v))

		if nowRow == numRows - 1 && !uping {
			uping = true
		} else if nowRow == 0 && uping {
			uping = false
		} else {}

		if uping {
			nowRow--
		} else {
			nowRow++
		}
	}

	res := make([]byte, 0)
	for _, v := range strings {
		res = append(res, v...)
	}
	return string(res)
}
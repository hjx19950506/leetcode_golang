package isNumber

import "strings"

//空格只能出现在首尾，所以开始先trim掉
//e只能出现一次，且后面要有数字
//.只能出现一次，且必须在e之前
//正负号只能出现在：1、数字最前面；2、e后面
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	e, dot, numBeforeE, numAfterE := false, false, false, true
	for i, v := range s {
		switch {
		case v == 'e' || v == 'E':
			if e || !numBeforeE {
				return false
			}
			e = true
			numAfterE = false
		case v == '.':
			if dot || e {
				return false
			}
			dot = true
		case v == '-' || v == '+':
			if i != 0 && !(s[i-1] == 'e' || s[i-1] == 'E') {
				return false
			}
		case v >= '0' && v <= '9':
			if !e {
				numBeforeE = true
			} else {
				numAfterE = true
			}
		default:
			return false
		}
	}
	return numBeforeE && numAfterE
}
package restoreIpAddresses

import (
	"strconv"
	"strings"
)

//dfs
//两次剪枝：
//第一次剪枝，剩余字符串能足够分；
//第二次剪枝，当前子串是合法的。
func restoreIpAddresses(s string) []string {
	res := []string{}
	temp := []string{}
	dfs(s, &res, temp)
	return res
}

func dfs(s string, res *[]string, temp []string) {
	if s == "" {
		if len(temp) == 4 {
			*res = append(*res, strings.Join(temp, "."))
		}
		return
	}
	for i := 0; i < 3; i++ {
		restSLen := len(s) - i - 1
		restSeg := 3 - len(temp)
		//第一次剪枝，剩余字符串能足够分
		if restSLen >= restSeg && restSLen <= 3 * restSeg {
			//第二次剪枝，当前子串是合法的
			if num, err := strconv.ParseInt(s[0:i+1], 10, 32); err == nil && num >= 0 && num <= 255 {
				temp = append(temp, strconv.FormatInt(num, 10))
				dfs(s[i+1:], res, temp)
				temp = temp[:len(temp)-1]
			}
		}
		//不能有前导零
		if s[0] == '0' {
			break
		}
	}
}
package countAndSay

import "strconv"

func countAndSay(n int) string {
	res := "1"

	for i := 2; i <= n; i++ {
		temp := ""

		curNum := res[0]
		curNumCount := 1
		for j := 1; j < len(res); j++ {
			if res[j] == curNum {
				curNumCount++
			} else {
				temp += strconv.Itoa(curNumCount) + string(curNum)
				curNum = res[j]
				curNumCount = 1
			}
		}
		temp += strconv.Itoa(curNumCount) + string(curNum)
		res = temp
	}

	return res
}
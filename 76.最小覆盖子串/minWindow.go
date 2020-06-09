package minWindow

//双指针滑动窗口
//先遍历一次t，用哈希存储每个字符出现的次数
//然后用左右指针在s上实现滑动窗口，并记录实时哈希
//首先向左移动r，直到当前窗口内字符不小于t对应字符出现次数
//然后向左移动l，实时检查哈希记录并更新最短字符串的左右索引
func minWindow(s string, t string) string {
	sLen, tLen:= len(s), len(t)
	if sLen == 0 || tLen == 0 {
		return ""
	}
	tMap, windowMap := map[byte]int{}, map[byte]int{}
	for i := range t {
		tMap[t[i]]++
	}

	leftAns, rightAns := -1, sLen
	for l, r := 0, 1; r <= sLen; r++ {
		if tMap[s[r-1]] > 0 {
			windowMap[s[r-1]]++
		}
		for isValid(tMap, windowMap) && l < r {
			if r - l < rightAns - leftAns {
				leftAns = l
				rightAns = r
			}
			if _, ok := tMap[s[l]]; ok {
				windowMap[s[l]] -= 1
			}
			l++
		}
	}
	if leftAns == -1 {
		return ""
	}
	return s[leftAns:rightAns]
}

func isValid(tMap, windowMap map[byte]int) bool {
	for k, v := range tMap {
		if windowMap[k] < v {
			return false
		}
	}
	return true
}
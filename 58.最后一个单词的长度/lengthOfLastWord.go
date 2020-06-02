package lengthOfLastWord

func lengthOfLastWord(s string) int {
	res := 0
	sLen := len(s)
	if sLen == 0 {
		return res
	}

	for i, count := 0, 0; i <= sLen; i++ {
		if i == sLen {
			if count != 0 {
				res = count
			}
			break
		}
		if s[i] != ' ' {
			count++
		} else {
			if count != 0 {
				res = count
			}
			count = 0
		}
	}
	return res
}
package lengthOfLongestSubstring


func lengthOfLongestSubstring(s string) int {
	l := len(s)
	log := [256]bool{}
	res, tmpRes := 0, 0

	//双指针迭代
	for i, j := 0, -1; j < l - 1; {
		if log[s[j+1]] == false {
			j++
			log[s[j]] = true
			tmpRes++
			if tmpRes > res {
				res = tmpRes
			}
		} else {
			for log[s[j+1]] {
				log[s[i]] = false
				i++
				tmpRes--
			}
		}
	}

	return res
}

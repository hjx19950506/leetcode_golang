package longestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var i int
	Loop:
	for i = 0; i <= len(strs[0]); i++ {
		for _, v := range strs {
			if v == "" {
				return ""
			}
			if !strings.HasPrefix(v, strs[0][:i]) {
				break Loop
			}
		}
	}

	return strs[0][:i-1]
}
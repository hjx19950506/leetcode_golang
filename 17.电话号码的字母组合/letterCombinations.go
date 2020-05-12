package letterCombinations

var m [][]byte = [][]byte{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q' ,'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}
var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	backTrack(digits, "", 0)
	return res
}

func backTrack(digits, temp string, p int) {
	if p == len(digits) {
		res = append(res, temp)
		return
	}

	digitToChar := m[digits[p] - '2']
	for i := 0; i < len(digitToChar); i++ {
		backTrack(digits, temp+string(digitToChar[i]), p+1)
	}
}
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
	backTrack(digits, "", 0)
	return res
}

func backTrack(digits, temp string, p int) {
	if len(digits) == len(temp) {
		res
	}
}
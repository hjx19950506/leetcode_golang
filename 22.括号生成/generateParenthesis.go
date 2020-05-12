package generateParenthesis

//回溯算法
//用l记录还剩下的(数量, l>0
//用r记录还剩下的(数量, r>l
var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	backtrack("", n, n)
	return res
}

func backtrack(temp string, l, r int) {
	if l == 0 && r == 0 {
		res = append(res, temp)
		return
	}
	if l > 0 {
		backtrack(temp+"(", l-1, r)
	}
	if r > l {
		backtrack(temp+")", l, r-1)
	}
}
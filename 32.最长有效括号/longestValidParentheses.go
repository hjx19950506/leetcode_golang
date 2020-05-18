package longestValidParentheses


//遍历s，使用两个stack分别存储字符的索引和值
//当栈顶和当前元素分别为'('和')'时出栈
//遍历索引栈寻找最大的索引间隔即为有效括号字符串
func longestValidParentheses(s string) int {
	sLen := len(s)
	res := 0
	if sLen < 2 {
		return res
	}

	//遍历s来填充两个stack
	stackOfIndex := make([]int, 0)
	stackOfSymbol := make([]int32, 0)
	for k, v := range s {
		if len(stackOfSymbol) == 0 || stackOfSymbol[len(stackOfSymbol)-1] == ')' || stackOfSymbol[len(stackOfSymbol)-1] == v {
			stackOfIndex = append(stackOfIndex, k)
			stackOfSymbol = append(stackOfSymbol, v)
		} else {
			stackOfIndex = stackOfIndex[:len(stackOfIndex)-1]
			stackOfSymbol = stackOfSymbol[:len(stackOfSymbol)-1]
		}
	}
	//遍历索引栈寻找最大间隔值
	stackOfIndex = append(stackOfIndex, sLen)
	res = stackOfIndex[0]
	for i := 1; i < len(stackOfIndex); i++ {
		if stackOfIndex[i] - stackOfIndex[i-1] - 1 > res {
			res = stackOfIndex[i] - stackOfIndex[i-1] - 1
		}
	}
	return res
}
package isValid

//用切片模拟栈
//栈顶匹配当前字符时出栈
//否则入栈
//根据最后栈是否为空判断
func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, b := range s {
		temp := byte(b)
		if len(stack) == 0 {
			stack = append(stack, temp)
		} else {
			back := stack[len(stack)-1]
			if (back == '(' && temp == ')') ||
				(back == '[' && temp == ']') ||
				(back == '{' && temp == '}'){
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, temp)
			}
		}
	}
	return len(stack) == 0
}

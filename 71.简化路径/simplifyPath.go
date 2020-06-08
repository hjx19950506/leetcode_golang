package simplifyPath

//遍历路径的所有文件夹名
//在栈中存储这些文件夹名以模拟当前路径
//最后使用该栈所剩文件夹名生成路径
func simplifyPath(path string) string {
	path += "/"
	pathLen := len(path)
	stack := []string{}
	directName := ""
	for i := 0; i < pathLen; i++ {
		if path[i] == '/' {
			if directName != "" {
				if directName == ".." {
					if len(stack) > 0 {
						stack = stack[:len(stack)-1]
					}
				} else if directName != "." {
					stack = append(stack, directName)
				}
			}
			directName = ""
		} else {
			directName += string(path[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res := ""
	for _, v := range stack {
		res += "/" + v
	}

	return res
}
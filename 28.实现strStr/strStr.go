package strStr


//KMP算法
//首先要根据needle初始化一个最长公共真前后缀长度数组
//然后利用该数组的元素去更少地移动needle完成匹配
func strStr(haystack string, needle string) int {
	len1, len2 := len(haystack), len(needle)
	if len2 == 0 {
		return 0
	}
	next := getNext(needle)
	i, j := 0, 0
	for i < len1 && j < len2 {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len2 {
		return i - j
	}
	return -1
}

func getNext(pattern string) []int {
	p_len := len(pattern)
	i, j := 0, -1
	next := make([]int, p_len)
	next[0] = -1
	for i < p_len - 1 {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			next[i] = j

		} else {
			j = next[j]
		}
	}
	return next
}
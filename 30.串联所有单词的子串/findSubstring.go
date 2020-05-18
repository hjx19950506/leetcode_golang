package findSubstring

import "reflect"


//滑动窗口法
//用words初始化一个map记录每个单词出现的次数
//以一个单词长度为步长遍历s，并使用temp记录窗口中单词出现的次数
//比较m和temp得知当前窗口是否满足要求
func findSubstring(s string, words []string) []int {
	res := []int{}
	sLen := len(s)
	numOfWords := len(words)
	if s == "" || numOfWords == 0 || numOfWords * len(words[0]) > sLen {
		return res
	}
	wordLen := len(words[0])
	windowLen := wordLen * numOfWords
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}

	for i := 0; i < wordLen && i <= sLen - windowLen; i++ {
		temp := make(map[string]int)
		//先把窗填满，然后做一次比较
		for j := i; j < i + windowLen; j += wordLen {
			temp[s[j:j+wordLen]]++
		}
		if reflect.DeepEqual(m, temp) {
			res = append(res, i)
		}
		//每次移动一步，将窗前面那个单词的次数减一
		//新进窗的单词加一
		//然后比较
		for j := i + wordLen; j <= sLen - windowLen; j += wordLen {
			if temp[s[j-wordLen:j]]--; temp[s[j-wordLen:j]] == 0 {
				delete(temp, s[j-wordLen:j])
			}
			temp[s[j+windowLen-wordLen:j+windowLen]]++
			if reflect.DeepEqual(m, temp) {
				res = append(res, j)
			}
		}
	}
	return res
}

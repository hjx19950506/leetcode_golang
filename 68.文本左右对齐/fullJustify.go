package fullJustify

import "strings"

//先确定每一行最多能放那几个单词
//然后再重排这些单词，填充空格
//最后一行做特殊处理
func fullJustify(words []string, maxWidth int) []string {
	length := len(words)
	res := make([]string, 0)
	if length == 0 {
		return res
	}
	for i := 0; i < length; {
		row := make([]string, 0)
		row = append(row, words[i])
		rowLen := len(words[i])
		for j := i+1; j < length; j++ {
			if rowLen + 1 + len(words[j]) <= maxWidth {
				row = append(row, words[j])
				rowLen += 1 + len(words[j])
			} else {
				break
			}
		}
		i += len(row)
		if i == length {
			temp := row[0]
			for i := 1; i < len(row); i++ {
				temp += " " + row[i]
			}
			temp += strings.Repeat(" ", maxWidth - len(temp))
			res = append(res, temp)
		} else {
			res = append(res, reorder(row, rowLen, maxWidth))
		}
	}
	return res
}

func reorder(src []string, rowLen, maxWidth int) string {
	dst := src[0]
	srcLen := len(src)
	spaceCount := maxWidth - rowLen + srcLen - 1
	if srcLen == 1 {
		dst += strings.Repeat(" ", spaceCount)
		return dst
	}
	aveSpace := spaceCount / (srcLen - 1)
	plusOneSpaceCount := spaceCount - aveSpace * (srcLen - 1)
	for i := 1; i < srcLen; i++ {
		if i <= plusOneSpaceCount {
			dst += " "
		}
		dst += strings.Repeat(" ", aveSpace)
		dst += src[i]
	}
	return dst
}
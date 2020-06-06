package fullJustify

func fullJustify(words []string, maxWidth int) []string {
	length := len(words)
	res := make([]string, 0)
	if length == 0 {
		return res
	}
	rows := make([][]string, 0)
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
		rows = append(rows, dump(row))
		i += len(row)
	}
	return res
}

//todo rerank
func dump(src []string) string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}
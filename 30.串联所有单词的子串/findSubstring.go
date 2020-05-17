package findSubstring

func findSubstring(s string, words []string) []int {
	if s == "" || words == nil {
		return nil
	}
	res := []int{}
	sLen := len(s)
	wordLen := len(words[0])
	numOfWords := len(words)
	m := make(map[string]int, numOfWords)
	for _, v := range words {
		if _, ok := m[v]; ok {
			m[v]++;
		} else {
			m[v] = 0;
		}
	}

	for i := 0; i < wordLen; i++ {
		for j := i; j + numOfWords * wordLen <  sLen; j += wordLen {


		}

	}
	return res
}

package longestPalindrome

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	l, r := 0, 0

	//初始化dp表格
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = 1
	}
	for i := 0; i < length - 1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
			l, r = i, i+1
		}
	}

	//填表
	for j := 1; j < length; j++ {
		for i := 0; i < length - 1; i++ {
			if dp[i+1][j-1] > 0 && s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				if j - i > r - l {
					l, r = i, j
				}
			}

		}
	}

	return s[l:r+1]
}

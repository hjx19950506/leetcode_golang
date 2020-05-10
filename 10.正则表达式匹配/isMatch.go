package isMatch

func isMatch(s string, p string) bool {
	len1, len2 := len(s), len(p)
	dp := make([][]bool, len1+1)
	for i := range dp {
		dp[i] = make([]bool, len2+1)
	}

	dp[0][0] = true
	for i := 2; i <= len2; i++ {
		dp[0][i] = p[i-1] == '*' && dp[0][i-2]
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if p[j-1] == '.' || s[i-1] == p[j-1] {
				//p[j-1]匹配s[i-1]
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				//p[j-1]为'*'
				if p[j-2] == s[i-1] || p[j-2] == '.'{
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]

				} else {
					//*前字符与当前s字符不同且不为.
					//那么只能0次
					dp[i][j] = dp[i][j-2]
				}
			} else {
				//p[j-1]不匹配s[i-1]，无需操作即为false
			}
		}
	}

	return dp[len1][len2]
}
package isMatch


//dp
func isMatch(s string, p string) bool {
	if p == "*" {
		return true
	}
	sLen, pLen := len(s), len(p)
	dp := make([][]bool, pLen+1)
	for i := 0; i <= pLen; i++ {
		dp[i] = make([]bool, sLen+1)
	}
	dp[0][0] = true
	for i := 1; i < sLen; i++ {
		dp[0][i] = false
	}
	for i := 1; i < pLen; i++ {
		dp[i][0] = dp[i-1][0] && p[i-1] == '*'
	}
	for i := 1; i <= pLen; i++ {
		for j := 1; j <= sLen; j++ {
			if p[i-1] == '?' || s[j-1] == p[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i-1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[pLen][sLen]
}
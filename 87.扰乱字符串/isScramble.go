package isScramble

//dp[i][j][lgt]为s1从i开始,s2从j开始长度为len的字符串是否互为扰乱
//不交换或交换两种情况,
//要么s1左与s2左、s1右与s2右互为扰乱，dp[i][j][k] && dp[i+k][j+k][lgt-k]
//要么s1左与s2右、s1右与s2左互为扰乱，dp[i][j+lgt-k][k] && dp[i+k][j][lgt-k]
//k从小到大增长，只要任何k使得上两式之一成立，dp[i][j][lgt]即为true
func isScramble(s1 string, s2 string) bool {
	l, l1 := len(s1), len(s2)
	if l != l1 {
		return false
	}
	if l == 0 {
		return true
	}
	dp := make([][][]bool, l1)
	for i := 0; i < l; i++ {
		dp[i] = make([][]bool, l)
		for j := 0; j < l; j++ {
			dp[i][j] = make([]bool, l+1)
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
				dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for lgt := 2; lgt <= l; lgt++ {
		for i := 0; i <= l - lgt; i++ {
			for j := 0; j <= l - lgt; j++ {
				for k := 1; k < lgt; k++ {
					if dp[i][j][k] && dp[i+k][j+k][lgt-k] {
						dp[i][j][lgt] = true
						break
					}
					if dp[i][j+lgt-k][k] && dp[i+k][j][lgt-k] {
						dp[i][j][lgt] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][l]
}
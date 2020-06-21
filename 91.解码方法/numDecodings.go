package numDecodings

//dp
//1、当当前数字为0且前一数字不为1或2则非法
//2、当当前数字与前一数字可以组成11-19，21-26时，
//	说明到当前位置字符和前一个字符有两种解读，
//	其解码种数dp[i] = dp[i-1] + dp[i-2]；
//3、否则和上一个的子串一样。
func numDecodings(s string) int {
	res := 0
	sLen := len(s)
	if sLen == 0 || s[0] == '0'{
		return res
	}
	if sLen == 1 {
		return 1
	}
	dp := make([]int, sLen)
	dp[0] = 1
	if s[1] == '0' && s[0] != '1' && s[0] != '2'{
		return 0
	}
	if (s[0] == '1' && s[1] != '0') || (s[0] == '2' && s[1] > '0' && s[1] < '7') {
		dp[1] = 2
	} else  {
		dp[1] = 1
	}
	for i := 2; i < sLen; i++ {
		//当前位置是0的话就需要通过前一数字判断是否合法
		if s[i] == '0' && s[i-1] != '1' && s[i-1] != '2'{
			return 0
		}
		if (s[i-1] == '1' && s[i] != '0') || (s[i-1] == '2' && s[i] > '0' && s[i] < '7') {
			dp[i] = dp[i-1] + dp[i-2]
		} else if s[i] == '0' {
			dp[i] = dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[sLen-1]
}
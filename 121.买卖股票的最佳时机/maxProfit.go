package maxProfit

//dp
//使用i表示第几天
//使用01状态标记当前是否持有
//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
//dp[i][1] = max(dp[i-1][1], -prices[i])
func maxProfit(prices []int) int {
	days := len(prices)
	if days == 0 {
		return 0
	}
	dp := make([][]int, days)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < days; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[days-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
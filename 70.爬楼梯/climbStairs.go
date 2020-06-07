package climbStairs

//dp
//f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
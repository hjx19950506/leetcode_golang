package getRow

//组合数+dp
//杨辉三角的每个数都是组合数
//第n行的第k个数为C(n,k) = n! / (k! * (n - k)!)
//即(n*(n-1)*(n-2)*...*(n-k+1))/k!
//可以看出每个数都可由前一个数进一步计算得到
//故可以使用dp
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for k := 1; k <= rowIndex; k++ {
		res[k] = res[k-1] * (rowIndex - k + 1) / k
	}
	return res
}
/*
//res[k] = combination(rowIndex, k)
func combination(n, k int) int {
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - i + 1) / i
	}
	return res
}
*/
package myPow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := 1.0
	flag1, flag2 := 1, 1
	if n < 0 {
		flag2 = -1
		n = -n
	}
	if x < 0 && n%2==1 {
		flag1 = -1
		x = -x
	}
	for i := 0; 1<<i <= n; i++ {
		if (n & (1<<i)) != 0 {
			res *= x
		}
		x *= x
	}
	if flag1 == -1 {
		res = -res
	}
	if flag2 == -1 {
		res = 1 / res
	}
	return res
}
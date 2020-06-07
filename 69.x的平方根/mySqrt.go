package mySqrt

//牛顿迭代
//即需求解f(x)=x^2-C的零根
//牛顿迭代公式为：x1=x0-f(x0)/f'(x0)-->x1=x0-(x0^2-C)/(2*x0)-->x1=(x0+C/x0)/2
func mySqrt(x int) int {
	x0 := x
	for x0 * x0 > x {
		x0 = (x0 + x / x0) / 2
	}
	return x0
}
package maxSubArray

//分治
func maxSubArray(nums []int) int{
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	mid := numsLen / 2
	lMax := maxSubArray(nums[:mid])
	rMax := maxSubArray(nums[mid:])
	mlMax, mlTemp := nums[mid-1], nums[mid-1]
	for i := mid-2; i >= 0; i-- {
		mlTemp += nums[i]
		if mlTemp > mlMax {
			mlMax = mlTemp
		}
	}
	mrMax, mrTemp := nums[mid], nums[mid]
	for i := mid+1; i < numsLen; i++ {
		mrTemp += nums[i]
		if mrTemp > mrMax {
			mrMax = mrTemp
		}
	}
	mMax := mlMax + mrMax
	return max(lMax, rMax, mMax)
}

func max(x, y, z int) int {
	if y > x {
		x = y
	}
	if z > x {
		x = z
	}
	return x
}


/*
//dp
//用f(i)代表以第i个数结尾的「连续子数组的最大和」
//f(i) = f(i-1)<0 ? nums[i] : f(i-1)+nums[i]
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < numsLen; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
*/
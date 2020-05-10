package maxArea

func maxArea(height []int) int {
	length := len(height)
	if length == 0 || length == 1{
		return 0
	}
	res := (length - 1) * Min(height[0], height[length-1])

	for l, r := 0, length-1; l < r; {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		tmpArea := (r - l) * Min(height[l], height[r])
		if tmpArea > res {
			res = tmpArea
		}
	}

	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
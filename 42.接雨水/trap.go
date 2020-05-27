package trap


//单调递减序列
//保证栈底元素左边是没有比他高的元素的
//所以可以以他为分界
func trap(height []int) int {
	res := 0
	decreaseStack := make([]int, 0)
	//decreaseStack[0] = 0
	for k, v := range height {
		if len(decreaseStack) == 0 || v < height[decreaseStack[len(decreaseStack)-1]] {
			decreaseStack = append(decreaseStack, k)
		} else {
			for v >= height[decreaseStack[len(decreaseStack)-1]] {
				top := decreaseStack[len(decreaseStack)-1]
				decreaseStack = decreaseStack[:len(decreaseStack)-1]
				if len(decreaseStack) == 0 {
					break
				}
				length := k - decreaseStack[len(decreaseStack)-1] - 1
				width := minInt(height[decreaseStack[len(decreaseStack)-1]], v) - height[top]
				res += length * width

			}
			decreaseStack = append(decreaseStack, k)
		}
	}

	return res
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
package largestRectangleArea

//单调递增栈，一次遍历即可
//分别记录每个柱子左右第一个低于其高度的柱子坐标
//当有数据可以入栈时，栈顶元素即为当前入栈元素向左第一个高度低于它的柱子
//当有元素需要出栈时，当前尝试入栈的元素即为出栈元素向右第一个高度低于它的柱子
//此外，最后额外插入一个高度-1柱子以排空栈中元素
func largestRectangleArea(heights []int) int {
	heights = append(heights, -1)
	length := len(heights)
	res := 0
	indexStack := []int{}
	lBound, rBound := make([]int, length), make([]int, length)

	for i := 0; i < length; {
		if len(indexStack) == 0{
			lBound[i] = -1
			indexStack = append(indexStack, i)
			i++
		} else if heights[indexStack[len(indexStack)-1]] < heights[i] {
			lBound[i] = indexStack[len(indexStack)-1]
			indexStack = append(indexStack, i)
			i++
		} else {
			for len(indexStack) != 0 && heights[indexStack[len(indexStack)-1]] >= heights[i] {
				rBound[indexStack[len(indexStack)-1]] = i
				indexStack = indexStack[:len(indexStack)-1]
			}
		}
	}

	for i := 0; i < length; i++ {
		if (rBound[i] - lBound[i] - 1) * heights[i] > res {
			res = (rBound[i] - lBound[i] - 1) * heights[i]
		}
	}
	return res
}
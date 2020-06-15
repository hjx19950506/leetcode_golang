package maximalRectangle

//利用84的函数
//本题每一层一上作为84函数的入参
//遍历每一层即可
func maximalRectangle(matrix [][]byte) int {
	res := 0
	rowCount := len(matrix)
	if rowCount == 0 {
		return res
	}
	colCount := len(matrix[0])
	heights := make([]int, colCount)
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		res = maxInt(res, largestRectangleArea(heights))
	}
	return res
}

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

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
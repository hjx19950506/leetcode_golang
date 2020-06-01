package insert

//56题基础上
//遍历intervals时以及前后各做一次判断
//此时是否需要插入newInterval
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 || newInterval[0] <= intervals[0][0] {
		res = append(res, newInterval)
	}
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			if intervals[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		}
		if res[len(res)-1][1] >= newInterval[0] {
			if newInterval[1] > res[len(res)-1][1] {
				res[len(res)-1][1] = newInterval[1]
			}
		} else if i < len(intervals)-1 && res[len(res)-1][1] < newInterval[0] && intervals[i+1][0] >= newInterval[0] {
			res = append(res, newInterval)
		}
	}
	if res[len(res)-1][1] < newInterval[0] {
		res = append(res, newInterval)
	}
	return res
}
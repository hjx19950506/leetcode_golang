package searchRange

import "testing"

//searchRange(nums []int, target int) []int
func TestSearchRange(t *testing.T) {
	cases := []struct{
		nums 	[]int
		target 	int
		want 	[]int
	} {
		{
			[]int{5,7,7,8,8,10},
			8,
			[]int{3,4},
		},
		{
			[]int{5,7,7,8,8,10},
			6,
			[]int{-1,-1},
		},
	}
}
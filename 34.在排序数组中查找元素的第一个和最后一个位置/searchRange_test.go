package searchRange

import (
	"fmt"
	"reflect"
	"testing"
)

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
		{
			[]int{5,7,7,7,7,10},
			7,
			[]int{1,4},
		},
		{
			[]int{7,7,7,7,7,7},
			7,
			[]int{0,5},
		},
		{
			[]int{0},
			0,
			[]int{0,0},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := searchRange(v.nums, v.target)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v. want)
			}
		})
	}
}
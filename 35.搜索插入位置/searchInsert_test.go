package searchInsert

import (
	"fmt"
	"testing"
)

//searchInsert(nums []int, target int) int {
func TestSearchInsert(t *testing.T) {
	cases := []struct{
		nums 	[]int
		target 	int
		want 	int
	} {
		{
			[]int{1,3,5,6},
			5,
			2,
		},
		{
			[]int{1,3,5,6},
			2,
			1,
		},
		{
			[]int{1,3,5,6},
			6,
			3,
		},
		{
			[]int{1,3,5,6},
			7,
			4,
		},
		{
			[]int{1,3,5,6},
			-1,
			0,
		},
		{
			[]int{},
			0,
			0,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			got := searchInsert(v.nums, v.target)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
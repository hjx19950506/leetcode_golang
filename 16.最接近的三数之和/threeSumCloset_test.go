package threeSumClosest

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	cases := []struct{
		nums 	[]int
		target 	int
		want	int
	} {
		{
			[]int{-1,2,1,-4},
			1,
			2,
		},
		{
			[]int{-1,2,1,-4},
			-7,
			-4,
		},
		{
			[]int{1,1,1,1},
			0,
			3,
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			got := threeSumClosest(v.nums, v.target)
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
